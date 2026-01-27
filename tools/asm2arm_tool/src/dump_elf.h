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

#ifndef DUMP_ELF_H
#define DUMP_ELF_H

#include "cal_depth.h"
#include "mc_bundle.h"

#include "llvm/ADT/StringRef.h"
#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCInstPrinter.h"

namespace tool {

namespace asm2arm {

/**
 * @brief 函数范围结构
 *
 * 表示一个函数的地址范围和名称
 */
struct FuncRange {
  uint64_t StartAddr; ///< 函数起始地址
  uint64_t EndAddr;   ///< 函数结束地址（左闭右开）
  std::string Name;   ///< 函数名称
};

/**
 * @brief 转储ELF文件
 *
 * 从ELF文件中提取信息并生成相应的输出文件
 * @param OutputPath 输出文件路径
 * @param ElfPath ELF文件路径
 * @param Bundle MC 上下文捆绑
 * @param Package 包名
 * @param BaseName 基础文件名
 * @param DumpTextSize 转储的文本段大小
 * @param Mode 模式
 */
void DumpElf(const std::string &OutputPath, llvm::StringRef ElfPath,
             tool::mc::MCContextBundle &Bundle, const std::string &Package,
             const std::string &BaseName, uint64_t &DumpTextSize,
             const std::string &Mode);

/**
 * @brief 转储子例程
 *
 * 转储指定入口基本块的子例程信息
 * @param EntryBB 入口基本块
 * @param Package 包名
 * @param OutputPath 输出文件路径
 * @param BaseName 基础文件名
 * @param SPDelta 栈指针变化列表
 * @param Depth 栈深度列表
 * @param DumpTextSize 转储的文本段大小
 */
void DumpSubr(const BasicBlock &EntryBB, const std::string &Package,
              const std::string &OutputPath, const std::string &BaseName,
              const std::vector<std::pair<uint64_t, int64_t>> &SPDelta,
              const std::vector<int64_t> &Depth, uint64_t DumpTextSize);

/**
 * @brief 转储模板
 *
 * 从模板文件生成输出文件
 * @param TmplFile 模板文件路径
 * @param Package 包名
 * @param OutputPath 输出文件路径
 * @param BaseName 基础文件名
 */
void DumpTmpl(const std::string &TmplFile, const std::string &Package,
              const std::string &OutputPath, const std::string &BaseName);

} // end namespace asm2arm

} // end namespace tool

#endif // DUMP_ELF_H