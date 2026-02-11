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

#ifndef STREAMER_SL_H
#define STREAMER_SL_H

#include "go_func_parser.h"
#include "mc_bundle.h"

#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCCodeEmitter.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCELFStreamer.h"
#include "llvm/MC/MCObjectWriter.h"
#include <cstddef>

/**
 * @file streamer_SL.h
 * @brief SL 流处理器头文件
 *
 * 实现 Plan9 格式的代码生成和流处理，用于生成符合 Plan9 格式的汇编代码
 */

/// 代码、数据对齐，细节如何处理
///     1 arm64代码段肯定是4字节对齐
///     2 数据段在每个标签对应的数据大小后补齐至4字节对齐
///     3 .align N 或 .p2align N 可以直接丢掉

namespace tool {
namespace asm2arm {

/**
 * @brief Plan9 流处理器
 *
 * 继承自 llvm::MCELFStreamer，用于生成符合 Plan9 格式的汇编代码
 */
class SLStreamer : public llvm::MCELFStreamer {
public:
  /**
   * @brief 构造函数
   *
   * @param Context MC上下文
   * @param AsmBackend MC汇编后端
   * @param ObjectWriter MC对象写入器
   * @param Emitter MC代码发射器
   * @param Out 输出流
   * @param Bundle MC上下文捆绑
   * @param BaseName 基础名称
   */
  SLStreamer(llvm::MCContext &Context,
             std::unique_ptr<llvm::MCAsmBackend> AsmBackend,
             std::unique_ptr<llvm::MCObjectWriter> ObjectWriter,
             std::unique_ptr<llvm::MCCodeEmitter> Emitter,
             llvm::raw_fd_ostream &Out, tool::mc::MCContextBundle &Bundle,
             const std::string &BaseName);

  /**
   * @brief 发射标签
   *
   * @param Sym 符号
   * @param Loc 源码位置
   */
  void emitLabel(llvm::MCSymbol *Sym, llvm::SMLoc Loc = {}) override;

  /**
   * @brief 发射指令
   *
   * @param Inst 指令
   * @param STI 子目标信息
   */
  void emitInstruction(const llvm::MCInst &Inst,
                       const llvm::MCSubtargetInfo &STI) override;

  /**
   * @brief 发射标识符
   *
   * .ident 不用处理
   * @param IdentString 标识符字符串
   */
  void emitIdent(llvm::StringRef IdentString) override;

  /**
   * @brief 发射整数值
   *
   * .byte | .short | .hword | .word | .quad | .dword | .xword
   * @param Value 值
   * @param Size 大小
   */
  void emitIntValue(uint64_t Value, unsigned Size) override;

  /**
   * @brief 发射填充
   *
   * .zero N |
   * @param NumBytes 字节数
   * @param FillValue 填充值
   * @param Loc 源码位置
   */
  void emitFill(const llvm::MCExpr &NumBytes, uint64_t FillValue,
                llvm::SMLoc Loc = llvm::SMLoc()) override;

  /**
   * @brief 发射字节
   *
   * .asciz "string" | .ascii "string"
   * .asciz 会被拆开调用两次emitBytes，第二次是单独的0
   * @param Data 数据
   */
  void emitBytes(llvm::StringRef Data) override;

  /**
   * @brief 发射对齐值
   *
   * .align N 或 .p2align N 不用处理
   * @param Alignment 对齐
   * @param Value 值
   * @param ValueSize 值大小
   * @param MaxBytesToEmit 最大发射字节数
   */
  void emitValueToAlignment(llvm::Align Alignment, int64_t Value = 0,
                            unsigned ValueSize = 1,
                            unsigned MaxBytesToEmit = 0) override;

private:
  tool::mc::MCContextBundle &Bundle; ///< MC上下文捆绑
  size_t IsTopEmit = 0;              ///< 是否顶部发射
  llvm::raw_fd_ostream &Out;         ///< 输出流
  std::string WordData;              ///< 字数据
  uint64_t ProgramCounter = 0;       ///< 程序计数器
  uint64_t StartProgramCounter = 0;  ///< 程序入口计数器
  const std::string &BaseName;       ///< 基础名称

public:
  /**
   * @brief 获取程序入口计数器
   *
   * @return 程序入口计数器
   */
  uint64_t GetStartProgramCounter();

  /**
   * @brief 生成字数据
   *
   * @param Value 值
   * @param Size 大小
   * @param Repeat 重复次数
   */
  void MakeWordData(uint64_t Value, unsigned Size, unsigned Repeat = 1);

  /**
   * @brief 刷新待处理字节，对齐至4字节
   */
  void FlushPendingBytes();

  /**
   * @brief 生成分支指令
   *
   * @param Token 指令序列
   * @param InstStr 指令字符串
   */
  void MakeBranchInst(const std::vector<std::string> &Token,
                      const std::string &InstStr);

  /**
   * @brief 生成分支
   *
   * @param Token 指令序列
   * @param InstStr 指令字符串
   */
  void MakeBranch(const std::vector<std::string> &Token,
                  const std::string &InstStr);

  /**
   * @brief 生成比较分支
   *
   * @param Token 指令序列
   * @param InstStr 指令字符串
   * @return 是否成功
   */
  bool MakeCmpareBranch(const std::vector<std::string> &Token,
                        const std::string &InstStr);

  /**
   * @brief 完成处理
   */
  void finish();
};

/**
 * @brief 生成声明头部
 *
 * @param Out 输出流
 * @param BaseName 基础名称
 * @param MaxDepth 最大深度
 */
void DumpDeclareHead(llvm::raw_fd_ostream &Out, const std::string &BaseName,
                     int64_t MaxDepth);

/**
 * @brief 生成声明尾部
 *
 * @param Out 输出流
 * @param BaseName 基础名称
 * @param ParseRes 解析结果
 * @param MaxDepth 最大深度
 */
void DumpDeclareTail(llvm::raw_fd_ostream &Out, const std::string &BaseName,
                     tool::ParseResult &ParseRes, int64_t MaxDepth);

/**
 * @brief 生成subr文件（SL模式）
 *
 * @param OutputPath 输出路径
 * @param Package 包名
 * @param BaseName 基础名称
 * @param StartPC 起始程序计数器
 * @param MaxDepth 最大深度
 */
void DumpSubrSL(const std::string &OutputPath, const std::string &Package,
                const std::string &BaseName, uint64_t StartPC,
                int64_t MaxDepth);

} // namespace asm2arm
} // namespace tool

#endif // STREAMER_SL_H