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
#include "go_func_parser.h"
#include "mc_bundle.h"
#include "streamer_SL.h"
#include "streamer_JIT.h"
#include "utils.h"

#include "lld/Common/Driver.h"
#include "llvm/ADT/SmallVector.h"
#include "llvm/CodeGen/CommandFlags.h"
#include "llvm/IR/LLVMContext.h"
#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCAsmInfo.h"
#include "llvm/MC/MCCodeEmitter.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCDisassembler/MCDisassembler.h"
#include "llvm/MC/MCInstPrinter.h"
#include "llvm/MC/MCObjectFileInfo.h"
#include "llvm/MC/MCObjectStreamer.h"
#include "llvm/MC/MCParser/MCAsmParser.h"
#include "llvm/MC/MCParser/MCTargetAsmParser.h"
#include "llvm/MC/MCStreamer.h"
#include "llvm/MC/MCTargetOptions.h"
#include "llvm/MC/TargetRegistry.h"
#include "llvm/Support/CommandLine.h"
#include "llvm/Support/Debug.h"
#include "llvm/Support/MemoryBuffer.h"
#include "llvm/Support/SourceMgr.h"
#include "llvm/Support/TargetSelect.h"
#include "llvm/Support/VirtualFileSystem.h"
#include "llvm/Support/raw_ostream.h"
#include "llvm/TargetParser/Host.h"
#include "llvm/TargetParser/Triple.h"

#include <cstddef>
#include <memory>
#include <string>
#include <vector>

using namespace llvm;
#define DEBUG_TYPE "main"
LLD_HAS_DRIVER(elf)

// COMMON Options
cl::OptionCategory CommonCategory("Tool COMMON Options");
static cl::opt<bool> Debug("debug", cl::desc("Enable debug output"),
                           cl::init(false), cl::Optional,
                           cl::cat(CommonCategory));
static cl::opt<std::string> ModeOption("mode", cl::desc("tool mode {JIT | SL}"),
                                       cl::value_desc("tool-mode"),
                                       cl::Required, cl::cat(CommonCategory));
static cl::opt<std::string> SourceFileOption("source",
                                             cl::desc("input ASM file"),
                                             cl::value_desc("ASM-file-path"),
                                             cl::Required,
                                             cl::cat(CommonCategory));
static cl::opt<std::string>
    OutputPathOption("output", cl::desc("Output path of *.go files"),
                     cl::value_desc("output-path"), cl::Required,
                     cl::cat(CommonCategory));
static cl::opt<std::string> LinkerScript("link-ld", cl::desc("linker script"),
                                         cl::value_desc("link-ld-path"),
                                         cl::Required, cl::cat(CommonCategory));
static cl::opt<std::string> PackageOption(
    "package", cl::desc("The package to which the generated Go file belongs"),
    cl::value_desc("package-name"), cl::Required, cl::cat(CommonCategory));
static cl::opt<std::string> FeaturesOption("features",
                                           cl::desc("features like +sve,+aes"),
                                           cl::value_desc("features"),
                                           cl::Optional,
                                           cl::cat(CommonCategory));

// JIT Options
cl::OptionCategory JITCategory("Tool JIT Options");
static cl::opt<std::string> TemplateFileOption("tmpl",
                                               cl::desc("Tmpl file path"),
                                               cl::value_desc("tmpl-path"),
                                               cl::cat(JITCategory));

// STATIC-LINK Options
cl::OptionCategory StaticLinkCategory("Tool STATIC-LINK Options");
static cl::opt<std::string>
    GoProtoOption("goproto", cl::desc("The go file that declares go functions"),
                  cl::value_desc("go-proto-path"), cl::cat(StaticLinkCategory));

/**
 * @brief 检查命令行选项
 *
 * 验证命令行选项是否有效
 * @return 选项是否有效
 */
bool CheckModeOptions() {
  if (ModeOption.empty()) {
    outs() << "--mode is empty\n";
    return false;
  }
  if (ModeOption != "JIT" && ModeOption != "SL") {
    outs() << "--mode is invalid\n";
    return false;
  }
  if (SourceFileOption.empty()) {
    outs() << "--source is empty\n";
    return false;
  }
  if (OutputPathOption.empty()) {
    outs() << "--output is empty\n";
    return false;
  }
  if (LinkerScript.empty()) {
    outs() << "--link-ld is empty\n";
    return false;
  }
  if (PackageOption.empty()) {
    outs() << "--package is empty\n";
    return false;
  }
  if (ModeOption == "JIT") {
    if (TemplateFileOption.empty()) {
      outs() << "tmpl is empty\n";
      return false;
    }
  }
  if (ModeOption == "SL") {
    if (GoProtoOption.empty()) {
      outs() << "goproto is empty\n";
      return false;
    }
  }
  return true;
}

/**
 * @brief 主函数
 *
 * 工具的入口点，负责处理命令行选项、生成目标文件、链接ELF文件、分析栈深度等
 * @param argc 命令行参数数量
 * @param argv 命令行参数数组
 * @return 执行结果，0表示成功，非0表示失败
 */
int main(int argc, char **argv) {
  // 解析命令行选项
  cl::HideUnrelatedOptions(
      {&CommonCategory, &JITCategory, &StaticLinkCategory});
  cl::ParseCommandLineOptions(argc, argv, "Tools for Go assembly conversion\n");
  if (!CheckModeOptions()) {
    return 1;
  }

  // 启用调试输出
  if (Debug) {
    DebugFlag = true;
  }

  // 提取命令行选项值
  std::string Mode = ModeOption;
  std::string SourceFile = SourceFileOption;
  std::string OutputPath = OutputPathOption;
  std::string Package = PackageOption;
  std::string TemplateFile = TemplateFileOption;
  std::string GoProto = GoProtoOption;

  // 获取源文件名（不含扩展名）
  auto BaseName = tool::GetSourceName(SourceFile);
  if (BaseName.empty()) {
    return 1;
  }

  // 初始化LLVM目标和MC组件
  InitializeAllTargets();
  InitializeAllTargetMCs();
  InitializeAllAsmPrinters();
  InitializeAllAsmParsers();
  InitializeAllDisassemblers();

  // 创建AArch64目标的MC上下文捆绑
  Triple TheTriple("aarch64-linux-gnu");
  tool::mc::MCContextBundle Bundle(TheTriple, FeaturesOption);
  tool::asm2arm::FindSP(Bundle);

  // 生成目标文件（object padding）
  SmallString<256> ObjFile;
  sys::path::append(ObjFile, OutputPath, (Twine(BaseName) + ".o").str());
  {
    auto MBExp = MemoryBuffer::getFile(SourceFile);
    if (!MBExp) {
      outs() << "getFile failed\n";
      return 1;
    }
    std::unique_ptr<MemoryBuffer> MB = std::move(*MBExp);

    SourceMgr SrcMgr;
    SrcMgr.AddNewSourceBuffer(std::move(MB), SMLoc());

    MCContext Ctx(TheTriple, &Bundle.getAsmInfo(), &Bundle.getRegisterInfo(),
                  &Bundle.getSubtargetInfo(), &SrcMgr,
                  &Bundle.getTargetOptions());
    std::unique_ptr<MCObjectFileInfo> MOFI(
        Bundle.getTarget().createMCObjectFileInfo(Ctx, false));
    Ctx.setObjectFileInfo(MOFI.get());

    auto MAB =
        std::unique_ptr<MCAsmBackend>(Bundle.getTarget().createMCAsmBackend(
            Bundle.getSubtargetInfo(), Bundle.getRegisterInfo(),
            Bundle.getTargetOptions()));
    auto MCE = std::unique_ptr<MCCodeEmitter>(
        Bundle.getTarget().createMCCodeEmitter(Bundle.getInstrInfo(), Ctx));

    std::error_code EC;
    raw_fd_ostream Out(ObjFile, EC, sys::fs::OF_None);
    if (EC) {
      outs() << EC.message() << "\n";
      return 1;
    }

    auto MOW = MAB->createObjectWriter(Out);
    // 不padding，会导致hasDefOfPhysReg功能失效
    auto Streamer = std::make_unique<tool::asm2arm::JITStreamer>(
        Ctx, std::move(MAB), std::move(MOW), std::move(MCE), Bundle);
    Streamer->initSections(false, Bundle.getSubtargetInfo());

    std::unique_ptr<MCAsmParser> Parser(
        llvm::createMCAsmParser(SrcMgr, Ctx, *Streamer, Bundle.getAsmInfo()));
    std::unique_ptr<MCTargetAsmParser> TAP(Bundle.getTarget().createMCAsmParser(
        Bundle.getSubtargetInfo(), *Parser, Bundle.getInstrInfo(),
        Bundle.getTargetOptions()));
    Parser->setTargetParser(*TAP);
    if (Parser->Run(false)) {
      outs() << "asm parse failed\n";
      return 1;
    }
    Streamer->finish();
  }

  // 链接生成ELF文件，消除相对地址
  SmallString<256> ELFFile;
  sys::path::append(ELFFile, OutputPath, (Twine(BaseName) + ".elf").str());
  {
    // ld.lld默认会优化某些指令组合，如adrp+add在寻址相对举例小于1MB时会被优化成一条adr指令，多出来的地址变为unknown
    // 这会打断BB块的数据流，为了避免这种情况，保持链接后与输入一致，增加--no-relax
    // 或者对于JIT这种来说，代码段一般都不会大于1MB，可以直接在生成汇编时加上-mcmodel=tiny
    std::string LinkerScriptArg = "--script=" + LinkerScript;
    std::vector<const char *> Args = {
        "ld.lld",        ObjFile.c_str(),         "-o",
        ELFFile.c_str(), LinkerScriptArg.c_str(), "--no-relax"};
    static const lld::DriverDef Drivers[] = {{lld::Gnu, lld::elf::link}};
    // 内部会调用cl::ParseCommandLineOptions，导致前面的选项被清空
    lld::Result Res = lld::lldMain(Args, outs(), errs(), Drivers);
    if (Res.retCode != 0) {
      outs() << "ld.lld failed\n";
      return 1;
    }
  }

  // 转储ELF文件信息
  uint64_t DumpSize = 0;
  tool::asm2arm::DumpElf(OutputPath, ELFFile, Bundle, Package, BaseName,
                         DumpSize, Mode);

  // 划分基本块
  std::string &EntryBlockName = BaseName;
  std::vector<tool::asm2arm::BasicBlock> BasicBlocks;
  int EntryBlockIndex =
      tool::asm2arm::SplitBasicBlocks(Bundle, BasicBlocks, EntryBlockName);

  std::vector<std::vector<size_t>> CFG;
  std::vector<tool::asm2arm::BasicBlockSP> BasicBlockSPVec;
  std::vector<std::pair<uint64_t, int64_t>> SPDelta;

  // 构建控制流图
  tool::asm2arm::BuildCFG(Bundle, BasicBlocks, CFG);
  LLVM_DEBUG(
      if (tool::asm2arm::HasCycle(BasicBlocks, CFG)) { dbgs() << "存在环\n"; });

  // 计算栈指针变化和最大栈深度
  tool::asm2arm::CalcSPDelta(Bundle, BasicBlocks, BasicBlockSPVec, SPDelta);
  auto MaxSPDepth = tool::asm2arm::ComputeMaxSPDepth(CFG, BasicBlockSPVec);
  LLVM_DEBUG(dbgs() << "EntryBlockIndex = " << EntryBlockIndex << "\n";);

  // 计算每条指令的最大栈深度
  auto InstMaxSPDepth = tool::asm2arm::ComputeMaxSPDepthAtInst(
      EntryBlockIndex, BasicBlocks, CFG, BasicBlockSPVec, SPDelta);
  int64_t MaxDepth = 0;
  for (auto X : InstMaxSPDepth) {
    MaxDepth = std::max(MaxDepth, X);
  }
  LLVM_DEBUG(for (size_t I = 0; I < InstMaxSPDepth.size(); I++) {
    dbgs() << "第" << I << "个SP" << SPDelta[I].first
           << " 最大深度:" << InstMaxSPDepth[I] << "\n";
  });

  // 根据模式执行不同的操作
  if (Mode == "JIT") {
    // JIT模式：生成子例程信息和模板文件
    tool::asm2arm::DumpSubr(BasicBlocks[EntryBlockIndex], Package, OutputPath,
                            BaseName, SPDelta, InstMaxSPDepth, DumpSize);
    tool::asm2arm::DumpTmpl(TemplateFile, Package, OutputPath, BaseName);
  }

  if (Mode == "SL") {
    // 静态链接模式：生成汇编文件和子例程信息
    SmallString<256> StaticFile;
    sys::path::append(StaticFile, OutputPath,
                      (Twine(BaseName) + "_arm64.s").str());
    auto MBExp = MemoryBuffer::getFile(SourceFile);
    if (!MBExp) {
      outs() << "getFile failed\n";
      return 1;
    }
    std::unique_ptr<MemoryBuffer> MB = std::move(*MBExp);

    SourceMgr SrcMgr;
    SrcMgr.AddNewSourceBuffer(std::move(MB), SMLoc());

    MCContext Ctx(TheTriple, &Bundle.getAsmInfo(), &Bundle.getRegisterInfo(),
                  &Bundle.getSubtargetInfo(), &SrcMgr,
                  &Bundle.getTargetOptions());
    std::unique_ptr<MCObjectFileInfo> MOFI(
        Bundle.getTarget().createMCObjectFileInfo(Ctx, false));
    Ctx.setObjectFileInfo(MOFI.get());

    auto MAB =
        std::unique_ptr<MCAsmBackend>(Bundle.getTarget().createMCAsmBackend(
            Bundle.getSubtargetInfo(), Bundle.getRegisterInfo(),
            Bundle.getTargetOptions()));
    auto MCE = std::unique_ptr<MCCodeEmitter>(
        Bundle.getTarget().createMCCodeEmitter(Bundle.getInstrInfo(), Ctx));

    std::error_code EC;
    raw_fd_ostream Out(StaticFile, EC, sys::fs::OF_None);
    if (EC) {
      outs() << EC.message() << "\n";
      return 1;
    }

    // 生成汇编文件头部
    tool::asm2arm::DumpDeclareHead(Out, BaseName, MaxDepth);

    raw_null_ostream NullOS;
    auto MOW = MAB->createObjectWriter(NullOS);

    // 创建Plan9流生成器
    auto Streamer = std::make_unique<tool::asm2arm::SLStreamer>(
        Ctx, std::move(MAB), std::move(MOW), std::move(MCE), Out, Bundle,
        BaseName);
    Streamer->initSections(false, Bundle.getSubtargetInfo());

    // 解析汇编文件
    std::unique_ptr<MCAsmParser> Parser(
        llvm::createMCAsmParser(SrcMgr, Ctx, *Streamer, Bundle.getAsmInfo()));
    std::unique_ptr<MCTargetAsmParser> TAP(Bundle.getTarget().createMCAsmParser(
        Bundle.getSubtargetInfo(), *Parser, Bundle.getInstrInfo(),
        Bundle.getTargetOptions()));
    Parser->setTargetParser(*TAP);
    if (Parser->Run(false)) {
      outs() << "asm parse failed\n";
      return 1;
    }
    Streamer->finish();

    // 解析Go函数签名并分配寄存器
    auto ParseRes = tool::ParseGoFile(GoProto);
    tool::AllocateRegisters(ParseRes);
    if (!ParseRes.Success()) {
      outs() << "Error: " << ParseRes.Error << "\n";
      return 1;
    }

    LLVM_DEBUG(for (const auto &[name, sig] : ParseRes.Funcs) {
      dbgs() << "Function: " << name << "\n";
      for (const auto &arg : sig.Params) {
        dbgs() << "  ARG " << (arg.Name.empty() ? "_" : arg.Name) << " ("
               << arg.Type << ", " << arg.Size << "B): " << arg.CReg.Name
               << " (C)\n";
      }
      for (const auto &res : sig.Results) {
        dbgs() << "  RET " << (res.Name.empty() ? "_" : res.Name) << " ("
               << res.Type << ", " << res.Size << "B): " << res.CReg.Name
               << " (C)\n";
      }
    });

    // 生成汇编文件尾部和子例程信息
    tool::asm2arm::DumpDeclareTail(Out, BaseName, ParseRes, MaxDepth);
    tool::asm2arm::DumpSubrSL(OutputPath, Package, BaseName,
                              Streamer->GetStartProgramCounter(), MaxDepth);
  }

  outs() << "ALL DONE\n";
  return 0;
}