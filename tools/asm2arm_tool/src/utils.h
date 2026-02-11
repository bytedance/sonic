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

#ifndef UTILS_H
#define UTILS_H

#include "mc_bundle.h"

#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCCodeEmitter.h"
#include "llvm/MC/MCELFStreamer.h"

namespace tool {

namespace asm2arm {

/**
 * @brief 查找栈指针寄存器
 *
 * 在 AArch64 架构中查找栈指针寄存器（SP）
 * @param Bundle MC 上下文捆绑
 */
void FindSP(tool::mc::MCContextBundle &Bundle);

/**
 * @brief 打印 AArch64 寄存器表
 *
 * 打印 AArch64 架构的寄存器映射表
 */
void PrintAArch64RegTable();

/**
 * @brief 打印指令辅助信息
 *
 * 打印 MC 指令的详细信息，包括操作码、操作数等
 * @param Inst MC 指令
 * @param Bundle MC 上下文捆绑
 * @param Addr 指令地址
 */
void PrintInstHelper(const llvm::MCInst &Inst,
                     tool::mc::MCContextBundle &Bundle, uint64_t Addr);

} // end namespace asm2arm

/**
 * @brief 检查字符串是否以指定前缀开始
 *
 * @param Str 要检查的字符串
 * @param Prefix 前缀
 * @return 是否以指定前缀开始
 */
bool StartWith(std::string_view Str, std::string_view Prefix);

/**
 * @brief 从路径获取源文件名
 *
 * 从完整路径中提取文件名（不含扩展名）
 * @param Path 文件路径
 * @return 源文件名
 */
std::string GetSourceName(llvm::StringRef Path);

/**
 * @brief 分词指令字符串
 *
 * 将指令字符串分解为标记列表
 * @param InstStr 指令字符串
 * @return 标记列表
 */
std::vector<std::string> TokenizeInstruction(const std::string &InstStr);

/**
 * @brief 输出标签
 *
 * 输出格式化的标签
 * @param Out 输出流
 * @param Label 标签名
 * @return 输出流引用
 */
llvm::raw_fd_ostream &OutLabel(llvm::raw_fd_ostream &Out,
                               llvm::StringRef Label);

/**
 * @brief 输出构建标签
 *
 * 输出固定的构建标签内容
 * @param Out 输出流
 */
void OutBuildTag(llvm::raw_ostream *Out);

} // end namespace tool

#endif // UTILS_H