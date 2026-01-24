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

#ifndef CAL_DEPTH_H
#define CAL_DEPTH_H

#include "mc_bundle.h"

#include <cstdint>

namespace tool {

namespace asm2arm {

/**
 * @brief 基本块结构
 *
 * 表示程序中的一个基本块，包含起始地址、结束地址和对应的指令索引
 */
struct BasicBlock {
  uint64_t StartAddr; ///< 基本块起始地址
  uint64_t EndAddr;   ///< 基本块结束地址（左闭右闭）
  size_t FirstIdx;    ///< 基本块第一条指令在Text中的索引
  size_t LastIdx;     ///< 基本块最后一条指令在Text中的索引
};

/**
 * @brief 基本块栈指针信息
 *
 * 表示基本块内栈指针的变化情况
 */
struct BasicBlockSP {
  int64_t NetDelta = 0; ///< 基本块内栈指针的总变化（字节）
  int64_t Peak = 0;     ///< 基本块内栈指针的峰值变化（<= 0，表示栈深度）
};

/**
 * @brief 基本块前缀和信息
 *
 * 表示基本块内每条指令执行前的栈指针变化前缀和
 */
struct BasicBlockPrefix {
  std::vector<int64_t> PreInstSP; ///< 基本块内每条指令执行前的栈指针变化前缀和
};

/**
 * @brief 分割基本块
 *
 * 根据指令流分割出程序的基本块
 * @param Bundle MC 上下文捆绑
 * @param BasicBlocks 输出基本块列表
 * @param EntryBB 入口基本块名称
 * @return 入口基本块的索引，失败返回 -1
 */
int SplitBasicBlocks(tool::mc::MCContextBundle &Bundle,
                     std::vector<BasicBlock> &BasicBlocks,
                     const std::string &EntryBB);

/**
 * @brief 构建控制流图
 *
 * 根据基本块构建控制流图（CFG）
 * @param Bundle MC 上下文捆绑
 * @param BasicBlocks 基本块列表
 * @param CFG 输出控制流图，CFG[i] 表示基本块 i 的后继基本块索引列表
 */
void BuildCFG(tool::mc::MCContextBundle &Bundle,
              std::vector<BasicBlock> &BasicBlocks,
              std::vector<std::vector<size_t>> &CFG);

/**
 * @brief 检查控制流图是否有环
 *
 * @param BasicBlocks 基本块列表
 * @param CFG 控制流图
 * @return 是否有环
 */
bool HasCycle(const std::vector<BasicBlock> &BasicBlocks,
              const std::vector<std::vector<size_t>> &CFG);

/**
 * @brief 计算栈指针变化
 *
 * 计算每个基本块内栈指针的变化情况
 * @param Bundle MC 上下文捆绑
 * @param BasicBlocks 基本块列表
 * @param BasicBlockSPVec 输出基本块栈指针信息列表
 * @param SPDelta 输出栈指针变化列表，每个元素为 (地址, 变化量)
 */
void CalcSPDelta(tool::mc::MCContextBundle &Bundle,
                 const std::vector<BasicBlock> &BasicBlocks,
                 std::vector<BasicBlockSP> &BasicBlockSPVec,
                 std::vector<std::pair<uint64_t, int64_t>> &SPDelta);

/**
 * @brief 计算最大栈深度
 *
 * 计算每个基本块的最大栈深度
 * @param CFG 控制流图
 * @param BasicBlockSPVec 基本块栈指针信息列表
 * @return 每个基本块的最大栈深度列表
 */
std::vector<int64_t>
ComputeMaxSPDepth(const std::vector<std::vector<size_t>> &CFG,
                  const std::vector<BasicBlockSP> &BasicBlockSPVec);

/**
 * @brief 计算每条指令的最大栈深度
 *
 * 计算指定入口基本块可达的每条指令的最大栈深度
 * @param EntryBB 入口基本块索引
 * @param BasicBlocks 基本块列表
 * @param CFG 控制流图
 * @param BasicBlockSPVec 基本块栈指针信息列表
 * @param SPDelta 栈指针变化列表
 * @return 每条指令的最大栈深度列表
 */
std::vector<int64_t> ComputeMaxSPDepthAtInst(
    size_t EntryBB, const std::vector<BasicBlock> &BasicBlocks,
    const std::vector<std::vector<size_t>> &CFG,
    const std::vector<BasicBlockSP> &BasicBlockSPVec,
    const std::vector<std::pair<uint64_t, int64_t>> &SPDelta);

} // end namespace asm2arm

} // end namespace tool

#endif // CAL_DEPTH_H