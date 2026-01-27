/*
 * Copyright 2026 Huawei Technologies Co., Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include "cal_depth.h"
#include "dump_elf.h"
#include "utils.h"

#include "llvm/ADT/StringRef.h"
#include "llvm/MC/MCInst.h"
#include "llvm/MC/MCInstPrinter.h"
#include "llvm/Support/Debug.h"
#include "llvm/Support/raw_ostream.h"

#include <cstdint>
#include <map>
#include <regex>
#include <set>
#include <sys/auxv.h>

#ifdef __ARM_FEATURE_SVE
#include <arm_sve.h>
#endif

using namespace llvm;
#define DEBUG_TYPE "cal_depth"

namespace tool {

namespace asm2arm {

extern std::map<std::string, unsigned> AArch64RegTable;
extern std::vector<MCInst> Text;
extern std::vector<uint64_t> TextPC;
extern std::unordered_map<uint64_t, size_t> Addr2Idx;
extern std::vector<FuncRange> Funcs;
extern std::set<uint64_t> Rets;

int SplitBasicBlocks(tool::mc::MCContextBundle &Bundle,
                     std::vector<BasicBlock> &BasicBlocks,
                     const std::string &EntryBB) {
  BasicBlocks.clear();
  if (Text.empty()) {
    return -1;
  }

  // 收集leader地址
  std::vector<uint64_t> Leaders;
  uint64_t EntryBBAddr = UINT64_MAX;

  auto TextBegin = Funcs[0].StartAddr;
  auto TextEnd = Funcs.back().EndAddr;
  // 函数入口leader
  for (const auto &F : Funcs) {
    Leaders.push_back(F.StartAddr);
    if (F.Name == EntryBB) {
      EntryBBAddr = F.StartAddr;
    }
  }

  // 跳转指令leader
  for (size_t I = 0; I < Text.size(); I++) {
    auto &Inst = Text[I];
    auto Pc = TextPC[I];
    if (Pc >= TextEnd || Pc < TextBegin) {
      break;
    }

    const auto &Desc = Bundle.getInstrInfo().get(Inst.getOpcode());
    if (Desc.isUnconditionalBranch()) {
      // 只有target，没有fall-through
      int64_t Offset = Inst.getOperand(Inst.getNumOperands() - 1).getImm();
      uint64_t Addr = static_cast<int64_t>(Pc) + Offset * 4;
      if (Addr2Idx.find(Addr) != Addr2Idx.end()) {
        Leaders.push_back(Addr);
      }
      continue;
    }
    if (Desc.isConditionalBranch() || Desc.isCall()) {
      // 有target，有fall-throough
      if (I + 1 < Text.size()) {
        Leaders.push_back(TextPC[I + 1]);
      }
      int64_t Offset = Inst.getOperand(Inst.getNumOperands() - 1).getImm();
      uint64_t Addr = static_cast<int64_t>(Pc) + Offset * 4;
      if (Addr2Idx.find(Addr) != Addr2Idx.end()) {
        Leaders.push_back(Addr);
      }
      continue;
    }
    if (Desc.isReturn() || Desc.isTrap() || Desc.isBarrier()) {
      continue;
    }
  }

  std::sort(Leaders.begin(), Leaders.end());
  Leaders.erase(std::unique(Leaders.begin(), Leaders.end()), Leaders.end());

  // 切分基本块
  for (size_t K = 0; K + 1 < Leaders.size(); K++) {
    uint64_t StartAddr = Leaders[K];
    uint64_t EndAddr = Leaders[K + 1] - 4;
    size_t First = Addr2Idx[StartAddr];
    size_t Last = Addr2Idx[EndAddr];
    BasicBlocks.push_back({StartAddr, EndAddr, First, Last});
  }
  BasicBlocks.push_back({Leaders.back(), TextEnd - 4, Addr2Idx[Leaders.back()],
                         Addr2Idx[TextEnd - 4]});
  for (size_t I = 0; I < BasicBlocks.size(); I++) {
    if (BasicBlocks[I].StartAddr == EntryBBAddr) {
      return I;
    }
  }
  return -1;
}

/**
 * @brief 根据PC地址查找基本块
 *
 * 使用二分查找在基本块列表中查找包含指定PC地址的基本块
 * @param PC 要查找的PC地址
 * @param BBs 基本块列表
 * @return 基本块的索引，未找到返回 SIZE_MAX
 */
static size_t FindBasicBlock(uint64_t PC,
                             const std::vector<BasicBlock> &BasicBlocks) {
  auto Pos = std::lower_bound(
      BasicBlocks.begin(), BasicBlocks.end(), PC,
      [](const BasicBlock &BB, uint64_t Addr) { return BB.EndAddr < Addr; });
  if (Pos == BasicBlocks.end()) {
    return SIZE_MAX;
  }
  return std::distance(BasicBlocks.begin(), Pos);
}

void BuildCFG(tool::mc::MCContextBundle &Bundle,
              std::vector<BasicBlock> &BasicBlocks,
              std::vector<std::vector<size_t>> &CFG) {
  CFG.clear();
  CFG.resize(BasicBlocks.size());

  for (size_t B = 0; B < BasicBlocks.size(); B++) {
    const BasicBlock &BB = BasicBlocks[B];
    MCInst &Inst = Text[BB.LastIdx];
    uint64_t LastPC = TextPC[BB.LastIdx];

    const auto &Desc = Bundle.getInstrInfo().get(Inst.getOpcode());

    // fall-through
    bool HasFallThrough = !(Desc.isUnconditionalBranch() || Desc.isReturn() ||
                            Desc.isTrap() || Desc.isBarrier());
    if (HasFallThrough && B + 1 < BasicBlocks.size()) {
      CFG[B].push_back(B + 1);
    }
    // target
    if (Desc.isUnconditionalBranch() || Desc.isConditionalBranch() ||
        Desc.isCall()) {
      int64_t Offset = Inst.getOperand(Inst.getNumOperands() - 1).getImm();
      uint64_t Addr = static_cast<int64_t>(LastPC) + Offset * 4;
      size_t TBB = FindBasicBlock(Addr, BasicBlocks);
      if (TBB != SIZE_MAX && TBB != B) {
        CFG[B].push_back(TBB);
      }
    }
  }
}

/**
 * @brief 控制流图环检测的DFS辅助函数
 *
 * 使用深度优先搜索检测控制流图中是否存在环
 * @param BBs 基本块列表
 * @param CFG 控制流图
 * @param V 当前访问的基本块索引
 * @param Color 节点颜色标记（0: 未访问, 1: 正在访问, 2: 已访问）
 * @return 是否存在环
 */
bool HasCycleDFS(const std::vector<BasicBlock> &BasicBlocks,
                 const std::vector<std::vector<size_t>> &CFG, size_t V,
                 std::vector<uint8_t> &Color) {
  Color[V] = 1; // 标记为正在访问
  for (auto U : CFG[V]) {
    if (Color[U] == 0) {
      if (HasCycleDFS(BasicBlocks, CFG, U, Color)) {
        return true;
      }
    } else if (Color[U] == 1) {
      // 发现回边，存在环
      return true;
    }
  }
  Color[V] = 2; // 标记为已访问
  return false;
}

bool HasCycle(const std::vector<BasicBlock> &BasicBlocks,
              const std::vector<std::vector<size_t>> &CFG) {
  std::vector<uint8_t> Color(CFG.size(),
                             0); // 0: 未访问, 1: 正在访问, 2: 已访问
  for (size_t I = 0; I < CFG.size(); I++) {
    if (Color[I] == 0 && HasCycleDFS(BasicBlocks, CFG, I, Color)) {
      return true;
    }
  }
  return false;
}

/**
 * @brief 获取栈指针调整量
 *
 * 分析指令，计算栈指针的调整量
 * @param Inst 指令对象
 * @param PC 指令地址
 * @param Bundle MC 上下文捆绑
 * @param SPDelta 输出栈指针变化列表
 * @return 栈指针调整量（正数表示栈指针增加，负数表示栈指针减少）
 */
int64_t GetSPAdjust(const MCInst &Inst, uint64_t PC,
                    tool::mc::MCContextBundle &Bundle,
                    std::vector<std::pair<uint64_t, int64_t>> &SPDelta) {
  const auto &Desc = Bundle.getInstrInfo().get(Inst.getOpcode());
  if (Desc.isReturn()) {
    Rets.insert(PC);
    return 0;
  }
  if (!Desc.hasDefOfPhysReg(Inst, AArch64RegTable["SP"],
                            Bundle.getRegisterInfo())) {
    return 0;
  }

  std::string InstLine;
  {
    raw_string_ostream Rso(InstLine);
    Bundle.getInstPrinter().printInst(&Inst, PC, {}, Bundle.getSubtargetInfo(),
                                      Rso);
  }
  LLVM_DEBUG(dbgs() << InstLine << "\n";);

  static const std::regex Re(R"(\s+(\w+)\s+.*sp.*#\s*(-?\d+)\b)");
  std::smatch Match;
  if (!std::regex_search(InstLine, Match, Re)) {
    return 0;
  }
  std::string Mnem = Match[1].str();
  int64_t Bytes = std::stoll(Match[2].str());

  int64_t Def = 0;
  uint64_t VL = 1;
#ifdef __ARM_FEATURE_SVE
  if (getauxval(AT_HWCAP) & HWCAP_SVE) {
    VL = svcntb();
  }
#endif
  if (StartWith(Mnem, "sub")) {
    Def = -Bytes;
  } else if (StartWith(Mnem, "addvl")) {
    Def = Bytes * VL;
  } else if (StartWith(Mnem, "subvl")) {
    Def = -Bytes * VL;
  } else if (StartWith(Mnem, "add") || StartWith(Mnem, "st") ||
             StartWith(Mnem, "ld")) {
    Def = Bytes;
  }
  LLVM_DEBUG(dbgs() << "        " << Mnem << " " << Bytes << " " << Def
                    << "\n\n";);
  SPDelta.push_back({PC, Def});
  return Def;
}

void CalcSPDelta(tool::mc::MCContextBundle &Bundle,
                 const std::vector<BasicBlock> &BasicBlocks,
                 std::vector<BasicBlockSP> &BasicBlockSPVec,
                 std::vector<std::pair<uint64_t, int64_t>> &SPDelta) {
  BasicBlockSPVec.assign(BasicBlocks.size(), {});
  for (size_t b = 0; b < BasicBlocks.size(); b++) {
    const auto &BB = BasicBlocks[b];
    int64_t Cur = 0;    // 当前栈指针变化
    int64_t MinCur = 0; // 最小栈指针变化（峰值深度）
    for (size_t i = BB.FirstIdx; i <= BB.LastIdx; i++) {
      Cur += GetSPAdjust(Text[i], TextPC[i], Bundle, SPDelta);
      if (Cur < MinCur) {
        MinCur = Cur;
      }
    }
    BasicBlockSPVec[b].NetDelta = Cur; // 基本块内栈指针总变化
    BasicBlockSPVec[b].Peak = MinCur;  // 基本块内栈指针峰值变化
  }
}

std::vector<int64_t>
ComputeMaxSPDepth(const std::vector<std::vector<size_t>> &CFG,
                  const std::vector<BasicBlockSP> &BasicBlockSPVec) {
  size_t N = BasicBlockSPVec.size();
  std::vector<int64_t> DP(N);

  // 初始化DP数组为每个基本块的峰值深度
  for (size_t i = 0; i < N; i++) {
    DP[i] = BasicBlockSPVec[i].Peak;
  }

  bool Updated;
  size_t Round = 0;
  // 迭代求解，直到没有更新或达到最大轮数
  for (; Round < N; Round++) {
    Updated = false;
    for (size_t U = 0; U < N; U++) {
      int64_t BestFromSucc = INT64_MAX;
      // 遍历所有后继基本块
      for (size_t V : CFG[U]) {
        if (DP[V] != INT64_MAX) {
          // 计算从后继基本块V回推到U的栈深度
          int64_t Candidate = BasicBlockSPVec[U].NetDelta + DP[V];
          if (Candidate < BestFromSucc) {
            BestFromSucc = Candidate;
          }
        }
      }
      // 如果找到更好的候选值，更新DP[U]
      if (BestFromSucc != INT64_MAX) {
        int64_t NewDp = std::min(DP[U], BestFromSucc);
        if (NewDp < DP[U]) {
          DP[U] = NewDp;
          Updated = true;
        }
      }
    }
    if (!Updated) {
      break;
    }
  }
  LLVM_DEBUG(dbgs() << "round = " << Round << " updated = " << Updated
                    << "\n";);
  // 将栈深度转换为正值（绝对值）
  for (auto &x : DP) {
    x = -x;
  }
  return DP;
}

/**
 * @brief 计算前向最小栈指针偏移
 *
 * 从指定入口基本块开始，计算每个基本块入口处的最小栈指针偏移
 * @param entry_bb 入口基本块索引
 * @param CFG 控制流图
 * @param BBSPVec 基本块栈指针信息列表
 * @return 每个基本块入口处的最小栈指针偏移列表
 */
std::vector<int64_t>
ComputeForwardMinSP(size_t EntryBB, const std::vector<std::vector<size_t>> &CFG,
                    const std::vector<BasicBlockSP> &BasicBlockSPVec) {
  size_t N = BasicBlockSPVec.size();
  std::vector<int64_t> EntryMinSP(N, INT64_MAX);
  EntryMinSP[EntryBB] = 0; // 入口处 SP 偏移为 0

  bool Changed;
  // 迭代传播，直到没有更新
  do {
    Changed = false;
    for (size_t u = 0; u < N; ++u) {
      if (EntryMinSP[u] == INT64_MAX) {
        continue; // 跳过不可达的基本块
      }

      // u 执行完后的 SP = 入口 SP + NetDelta
      int64_t SpAfterU = EntryMinSP[u] + BasicBlockSPVec[u].NetDelta;

      // 传播到所有后继基本块
      for (size_t v : CFG[u]) {
        if (SpAfterU < EntryMinSP[v]) { // 更深（更小）
          EntryMinSP[v] = SpAfterU;
          Changed = true;
        }
      }
    }
  } while (Changed);

  return EntryMinSP;
}

/**
 * @brief 预处理基本块前缀和
 *
 * 计算每个基本块内每条指令执行前的栈指针变化前缀和
 * @param BBs 基本块列表
 * @param SPDeltaMap 栈指针变化映射（地址 -> 变化量）
 * @return 每个基本块的前缀和信息列表
 */
std::vector<BasicBlockPrefix>
PreprocessBBPrefixes(const std::vector<BasicBlock> &BasicBlocks,
                     const std::unordered_map<uint64_t, int64_t> &SPDeltaMap) {
  std::vector<BasicBlockPrefix> BasicBlockPreInstSP(BasicBlocks.size());

  for (size_t bb_id = 0; bb_id < BasicBlocks.size(); ++bb_id) {
    auto &bb = BasicBlocks[bb_id];
    size_t count = bb.LastIdx - bb.FirstIdx + 1;
    BasicBlockPreInstSP[bb_id].PreInstSP.resize(count, 0);

    int64_t cum = 0; // 累积栈指针变化
    for (size_t i = 0; i < count; ++i) {
      BasicBlockPreInstSP[bb_id].PreInstSP[i] =
          cum; // 记录指令执行前的栈指针变化
      uint64_t pc = TextPC[bb.FirstIdx + i];
      if (auto it = SPDeltaMap.find(pc); it != SPDeltaMap.end()) {
        cum += it->second; // 更新累积栈指针变化
      }
    }
  }

  return BasicBlockPreInstSP;
}

/**
 * @brief 计算每条指令的最大栈深度
 *
 * 根据基本块入口栈指针偏移和前缀和，计算每条指令的最大栈深度
 * @param SPDelta 栈指针变化列表
 * @param BBs 基本块列表
 * @param EntryMinSP 每个基本块入口处的最小栈指针偏移
 * @param BBPreInstSP 每个基本块的前缀和信息
 * @return 每条指令的最大栈深度列表（正值）
 */
std::vector<int64_t> CalculateMaxDepthAtInstructions(
    const std::vector<std::pair<uint64_t, int64_t>> &SPDelta,
    const std::vector<BasicBlock> &BasicBlocks,
    const std::vector<int64_t> &EntryMinSP,
    const std::vector<BasicBlockPrefix> &BasicBlockPreInstSP) {
  std::vector<int64_t> MaxDepthAtInst(SPDelta.size(), 0);

  for (size_t i = 0; i < SPDelta.size(); ++i) {
    uint64_t addr = SPDelta[i].first;
    auto it_idx = Addr2Idx.find(addr);
    if (it_idx == Addr2Idx.end()) {
      continue;
    }
    size_t inst_idx = it_idx->second;

    // 找到所属基本块
    size_t bb_id = SIZE_MAX;
    for (size_t b = 0; b < BasicBlocks.size(); ++b) {
      if (inst_idx >= BasicBlocks[b].FirstIdx &&
          inst_idx <= BasicBlocks[b].LastIdx) {
        bb_id = b;
        break;
      }
    }
    if (bb_id == SIZE_MAX) {
      continue;
    }

    // 如果该基本块不可达（从 EntryBB 出发）
    if (EntryMinSP[bb_id] == INT64_MAX) {
      MaxDepthAtInst[i] = 0; // 不可达
      continue;
    }

    // 计算指令在基本块内的偏移
    size_t offset_in_bb = inst_idx - BasicBlocks[bb_id].FirstIdx;
    int64_t prefix = BasicBlockPreInstSP[bb_id].PreInstSP[offset_in_bb];

    // 到达该指令时的 SP 偏移 = 基本块入口 SP + 基本块内前缀
    int64_t sp_at_inst = EntryMinSP[bb_id] + prefix;

    // 最大深度（正值）
    MaxDepthAtInst[i] = -sp_at_inst;
  }

  return MaxDepthAtInst;
}

std::vector<int64_t> ComputeMaxSPDepthAtInst(
    size_t EntryBB, const std::vector<BasicBlock> &BasicBlocks,
    const std::vector<std::vector<size_t>> &CFG,
    const std::vector<BasicBlockSP> &BasicBlockSPVec,
    const std::vector<std::pair<uint64_t, int64_t>> &SPDelta) {
  // Step 1: 计算从指定入口基本块出发的 EntryMinSP
  auto EntryMinSP = ComputeForwardMinSP(EntryBB, CFG, BasicBlockSPVec);

  // 构建 SPDeltaMap
  std::unordered_map<uint64_t, int64_t> SPDeltaMap(SPDelta.begin(),
                                                   SPDelta.end());

  // Step 2: 预处理每个基本块的 PreInstSP
  auto BasicBlockPreInstSP = PreprocessBBPrefixes(BasicBlocks, SPDeltaMap);

  // Step 3: 结合 1 和 2 的结果计算从指定入口到达每个基本块
  // 块的每条指令的可能最大深度
  auto MaxDepths = CalculateMaxDepthAtInstructions(
      SPDelta, BasicBlocks, EntryMinSP, BasicBlockPreInstSP);

  return MaxDepths;
}

} // end namespace asm2arm

} // end namespace tool