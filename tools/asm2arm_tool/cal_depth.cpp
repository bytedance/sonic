#include "cal_depth.h"
#include "dump_elf.h"
#include "utils.h"

#include "llvm/ADT/StringRef.h"
#include "llvm/MC/MCInst.h"
#include "llvm/MC/MCInstPrinter.h"
#include "llvm/Support/raw_ostream.h"

#include <cstdint>
#include <map>
#include <regex>

#ifdef __ARM_FEATURE_SVE
#include <arm_sve.h>
#endif

using namespace llvm;

extern std::map<std::string, unsigned> AArch64RegTable;
extern std::vector<MCInst> Text;
extern std::vector<uint64_t> TextPC;
extern std::unordered_map<uint64_t, size_t> Addr2Idx;
extern std::vector<FuncRange> Funcs;

int SplitBasicBlocks(MCContextBundle &Bundle, std::vector<BasicBlock> &BBs, const std::string &EntryBB)
{
    BBs.clear();
    if (Text.empty()) {
        return -1;
    }

    // 收集leader地址
    std::vector<uint64_t> Leaders;
    uint64_t EntryBBAddr = UINT64_MAX;

    auto TextEnd = Funcs.back().EndAddr;
    // 函数入口leader
    for (const auto &F : Funcs) {
        Leaders.push_back(F.StartAddr);
        if (F.Name == EntryBB) {
            EntryBBAddr = F.StartAddr;
        }
    }

    // 跳转指令leader
    for (size_t i = 0; i < Text.size(); i++) {
        auto &Inst = Text[i];
        auto Pc = TextPC[i];
        if (Pc > TextEnd) {
            break;
        }

        const auto &Desc = Bundle.getMCInstrInfo().get(Inst.getOpcode());
        if (Desc.isUnconditionalBranch()) {
            // 只有target，没有fall-through
            int64_t Offset = Inst.getOperand(Inst.getNumOperands() - 1).getImm();
            uint64_t Addr = static_cast<int64_t>(Pc) + Offset * 4;
            Leaders.push_back(Addr);
            continue;
        }
        if (Desc.isConditionalBranch() || Desc.isCall()) {
            // 有target，有fall-throough
            if (i + 1 < Text.size()) {
                Leaders.push_back(TextPC[i + 1]);
            }
            int64_t Offset = Inst.getOperand(Inst.getNumOperands() - 1).getImm();
            uint64_t Addr = static_cast<int64_t>(Pc) + Offset * 4;
            Leaders.push_back(Addr);
            continue;
        }
        if (Desc.isReturn() || Desc.isTrap() || Desc.isBarrier()) {
            continue;
        }
    }

    std::sort(Leaders.begin(), Leaders.end());
    Leaders.erase(std::unique(Leaders.begin(), Leaders.end()), Leaders.end());

    // 切分BB
    for (size_t k = 0; k + 1 < Leaders.size(); k++) {
        uint64_t SAddr = Leaders[k];
        uint64_t EAddr = Leaders[k + 1] - 4;
        size_t First = Addr2Idx[SAddr];
        size_t Last = Addr2Idx[EAddr];
        BBs.push_back({SAddr, EAddr, First, Last});
    }
    BBs.push_back({Leaders.back(), TextEnd, Addr2Idx[Leaders.back()], Addr2Idx[TextEnd]});
    for (size_t i = 0; i < BBs.size(); i++) {
        if (BBs[i].StartAddr == EntryBBAddr) {
            return i;
        }
    }
    return -1;
}

static size_t FindBB(uint64_t PC, const std::vector<BasicBlock> &BBs)
{
    auto Pos = std::lower_bound(
        BBs.begin(), BBs.end(), PC, [](const BasicBlock &BB, uint64_t Addr) { return BB.EndAddr < Addr; });
    if (Pos == BBs.end()) {
        return SIZE_MAX;
    }
    return std::distance(BBs.begin(), Pos);
}

void BuildCFG(MCContextBundle &Bundle, std::vector<BasicBlock> &BBs, std::vector<std::vector<size_t>> &CFG)
{
    CFG.clear();
    CFG.resize(BBs.size());

    for (size_t b = 0; b < BBs.size(); b++) {
        const BasicBlock &BB = BBs[b];
        MCInst &Inst = Text[BB.LastIdx];
        uint64_t LastPC = TextPC[BB.LastIdx];

        const auto &Desc = Bundle.getMCInstrInfo().get(Inst.getOpcode());

        // fall-through
        bool HasFallThrough = !(Desc.isUnconditionalBranch() || Desc.isReturn() || Desc.isTrap() || Desc.isBarrier());
        if (HasFallThrough && b + 1 < BBs.size()) {
            CFG[b].push_back(b + 1);
        }
        // target
        if (Desc.isUnconditionalBranch() || Desc.isConditionalBranch() || Desc.isCall()) {
            int64_t Offset = Inst.getOperand(Inst.getNumOperands() - 1).getImm();
            uint64_t Addr = static_cast<int64_t>(LastPC) + Offset * 4;
            size_t TBB = FindBB(Addr, BBs);
            if (TBB != SIZE_MAX && TBB != b) {
                CFG[b].push_back(TBB);
            }
        }
    }
}

std::vector<uint64_t> Cycle;
bool HasCycleDFS(const std::vector<BasicBlock> &BBs, const std::vector<std::vector<size_t>> &CFG, size_t V,
    std::vector<uint8_t> &Color)
{
    Cycle.push_back(BBs[V].EndAddr);
    Color[V] = 1;
    for (auto U : CFG[V]) {
        if (Color[U] == 0) {
            if (HasCycleDFS(BBs, CFG, U, Color)) {
                return true;
            }
        } else if (Color[U] == 1) {
            return true;
        }
    }
    Color[V] = 2;
    Cycle.pop_back();
    return false;
}

bool HasCycle(const std::vector<BasicBlock> &BBs, const std::vector<std::vector<size_t>> &CFG)
{
    std::vector<uint8_t> Color(CFG.size(), 0);
    for (size_t i = 0; i < CFG.size(); i++) {
        if (Color[i] == 0 && HasCycleDFS(BBs, CFG, i, Color)) {
            return true;
        }
    }
    return false;
}

int64_t GetSPAdjust(
    const MCInst &Inst, uint64_t PC, MCContextBundle &Bundle, std::vector<std::pair<uint64_t, int64_t>> &SPDelta)
{
    const auto &Desc = Bundle.getMCInstrInfo().get(Inst.getOpcode());
    if (!Desc.hasDefOfPhysReg(Inst, AArch64RegTable["SP"], Bundle.getMCRegisterInfo())) {
        return 0;
    }

    std::string InstLine;
    {
        raw_string_ostream Rso(InstLine);
        Bundle.getMCInstPrinter().printInst(&Inst, PC, {}, Bundle.getMCSubtargetInfo(), Rso);
    }
    outs() << InstLine << "\n";

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
    VL = svcntb();
#endif
    if (StartWith(Mnem, "sub")) {
        Def = -Bytes;
    } else if (StartWith(Mnem, "addvl")) {
        Def = Bytes * VL;
    } else if (StartWith(Mnem, "subvl")) {
        Def = -Bytes * VL;
    } else if (StartWith(Mnem, "add") || StartWith(Mnem, "st") || StartWith(Mnem, "ld")) {
        Def = Bytes;
    }
    outs() << "        " << Mnem << " " << Bytes << " " << Def << "\n\n";
    SPDelta.push_back({PC, Def});
    return Def;
}

void CalcSPDelta(MCContextBundle &Bundle, const std::vector<BasicBlock> &BBs, std::vector<BBSP> &BBSPVec,
    std::vector<std::pair<uint64_t, int64_t>> &SPDelta)
{
    BBSPVec.assign(BBs.size(), {});
    for (size_t b = 0; b < BBs.size(); b++) {
        const auto &BB = BBs[b];
        int64_t Cur = 0;
        int64_t MinCur = 0;
        for (size_t i = BB.FirstIdx; i <= BB.LastIdx; i++) {
            Cur += GetSPAdjust(Text[i], TextPC[i], Bundle, SPDelta);
            if (Cur < MinCur) {
                MinCur = Cur;
            }
        }
        BBSPVec[b].NetDelta = Cur;
        BBSPVec[b].Peak = MinCur;
    }
}

std::vector<int64_t> ComputeMaxSPDepth(const std::vector<std::vector<size_t>> &CFG, const std::vector<BBSP> &BBSPVec)
{
    size_t N = BBSPVec.size();
    std::vector<int64_t> DP(N);

    for (size_t i = 0; i < N; i++) {
        DP[i] = BBSPVec[i].Peak;
    }

    bool Updated;
    size_t Round = 0;
    for (; Round < N; Round++) {
        Updated = false;
        for (size_t U = 0; U < N; U++) {
            int64_t BestFromSucc = INT64_MAX;
            for (size_t V : CFG[U]) {
                if (DP[V] != INT64_MAX) {
                    int64_t Candidata = BBSPVec[U].NetDelta + DP[V];
                    if (Candidata < BestFromSucc) {
                        BestFromSucc = Candidata;
                    }
                }
            }
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
    outs() << "round = " << Round << " updated = " << Updated << "\n";
    for (auto &x : DP) {
        x = -x;
    }
    return DP;
}

std::vector<int64_t> ComputeForwardMinSP(
    size_t entry_bb, const std::vector<std::vector<size_t>> &CFG, const std::vector<BBSP> &BBSPVec)
{
    size_t N = BBSPVec.size();
    std::vector<int64_t> EntryMinSP(N, INT64_MAX);
    EntryMinSP[entry_bb] = 0;  // 入口处 SP 偏移为 0

    bool changed;
    do {
        changed = false;
        for (size_t u = 0; u < N; ++u) {
            if (EntryMinSP[u] == INT64_MAX) {
                continue;
            }

            // u 执行完后的 SP = 入口 SP + NetDelta
            int64_t sp_after_u = EntryMinSP[u] + BBSPVec[u].NetDelta;

            // 传播到所有后继
            for (size_t v : CFG[u]) {
                if (sp_after_u < EntryMinSP[v]) {  // 更深（更小）
                    EntryMinSP[v] = sp_after_u;
                    changed = true;
                }
            }
        }
    } while (changed);

    return EntryMinSP;
}

std::vector<BBPrefix> PreprocessBBPrefixes(
    const std::vector<BasicBlock> &BBs, const std::unordered_map<uint64_t, int64_t> &SPDeltaMap)
{
    std::vector<BBPrefix> BBPreInstSP(BBs.size());

    for (size_t bb_id = 0; bb_id < BBs.size(); ++bb_id) {
        auto &bb = BBs[bb_id];
        size_t count = bb.LastIdx - bb.FirstIdx + 1;
        BBPreInstSP[bb_id].PreInstSP.resize(count, 0);

        int64_t cum = 0;
        for (size_t i = 0; i < count; ++i) {
            BBPreInstSP[bb_id].PreInstSP[i] = cum;
            uint64_t pc = TextPC[bb.FirstIdx + i];
            if (auto it = SPDeltaMap.find(pc); it != SPDeltaMap.end()) {
                cum += it->second;
            }
        }
    }

    return BBPreInstSP;
}

std::vector<int64_t> CalculateMaxDepthAtInstructions(const std::vector<std::pair<uint64_t, int64_t>> &SPDelta,
    const std::vector<BasicBlock> &BBs, const std::vector<int64_t> &EntryMinSP,
    const std::vector<BBPrefix> &BBPreInstSP)
{
    std::vector<int64_t> MaxDepthAtInst(SPDelta.size(), 0);

    for (size_t i = 0; i < SPDelta.size(); ++i) {
        uint64_t addr = SPDelta[i].first;
        auto it_idx = Addr2Idx.find(addr);
        if (it_idx == Addr2Idx.end()) {
            continue;
        }
        size_t inst_idx = it_idx->second;

        // 找到所属 BB
        size_t bb_id = SIZE_MAX;
        for (size_t b = 0; b < BBs.size(); ++b) {
            if (inst_idx >= BBs[b].FirstIdx && inst_idx <= BBs[b].LastIdx) {
                bb_id = b;
                break;
            }
        }
        if (bb_id == SIZE_MAX) {
            continue;
        }

        // 如果该 BB 不可达（从 entry_bb 出发）
        if (EntryMinSP[bb_id] == INT64_MAX) {
            MaxDepthAtInst[i] = 0;  // 不可达
            continue;
        }

        // 计算指令在 BB 内的偏移
        size_t offset_in_bb = inst_idx - BBs[bb_id].FirstIdx;
        int64_t prefix = BBPreInstSP[bb_id].PreInstSP[offset_in_bb];

        // 到达该指令时的 SP 偏移 = BB 入口 SP + BB 内前缀
        int64_t sp_at_inst = EntryMinSP[bb_id] + prefix;

        // 最大深度（正值）
        MaxDepthAtInst[i] = -sp_at_inst;
    }

    return MaxDepthAtInst;
}

std::vector<int64_t> ComputeMaxSPDepthAtInst(size_t entry_bb, const std::vector<BasicBlock> &BBs,
    const std::vector<std::vector<size_t>> &CFG, const std::vector<BBSP> &BBSPVec,
    const std::vector<std::pair<uint64_t, int64_t>> &SPDelta)
{
    // Step 1: 计算从指定入口 BB 出发的 EntryMinSP
    auto EntryMinSP = ComputeForwardMinSP(entry_bb, CFG, BBSPVec);

    // 构建 SPDeltaMap
    std::unordered_map<uint64_t, int64_t> SPDeltaMap(SPDelta.begin(), SPDelta.end());

    // Step 2: 预处理每个 BB 的 PreInstSP
    auto BBPreInstSP = PreprocessBBPrefixes(BBs, SPDeltaMap);

    // Step 3: 结合 1 和 2 的结果计算从指定入口到达每个 BB 块的每条指令的可能最大深度
    auto maxDepths = CalculateMaxDepthAtInstructions(SPDelta, BBs, EntryMinSP, BBPreInstSP);

    return maxDepths;
}