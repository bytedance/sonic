#include "mc_bundle.h"
#include "streamer.h"
#include "plan9_streamer.h"
#include "dump_elf.h"
#include "cal_depth.h"
#include "go_func_parser.h"
#include "utils.h"

#include "llvm/ADT/SmallVector.h"
#include "llvm/CodeGen/CommandFlags.h"
#include "llvm/IR/LLVMContext.h"
#include "llvm/MC/MCDisassembler/MCDisassembler.h"
#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCAsmInfo.h"
#include "llvm/MC/MCCodeEmitter.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCInstPrinter.h"
#include "llvm/MC/MCObjectFileInfo.h"
#include "llvm/MC/MCParser/MCAsmParser.h"
#include "llvm/MC/MCParser/MCTargetAsmParser.h"
#include "llvm/MC/MCStreamer.h"
#include "llvm/MC/MCObjectStreamer.h"
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
#include "lld/Common/Driver.h"

#include <cstddef>
#include <memory>
#include <string>
#include <vector>

using namespace llvm;
#define DEBUG_TYPE "main"
LLD_HAS_DRIVER(elf)

// COMMON Options
cl::OptionCategory CommonCategory("Tool COMMON Options");
static cl::opt<bool> Debug(
    "debug", cl::desc("Enable debug output"), cl::init(false), cl::Optional, cl::cat(CommonCategory));
static cl::opt<std::string> Mode_(
    "mode", cl::desc("tool mode {JIT | SL}"), cl::value_desc("tool-mode"), cl::Required, cl::cat(CommonCategory));
static cl::opt<std::string> SourceFile_(
    "source", cl::desc("input ASM file"), cl::value_desc("ASM-file-path"), cl::Required, cl::cat(CommonCategory));
static cl::opt<std::string> OutputPath_("output", cl::desc("Output path of *.go files"), cl::value_desc("output-path"),
    cl::Required, cl::cat(CommonCategory));
static cl::opt<std::string> LdScript(
    "link-ld", cl::desc("linker script"), cl::value_desc("link-ld-path"), cl::Required, cl::cat(CommonCategory));
static cl::opt<std::string> Package_("package", cl::desc("The package to which the generated Go file belongs"),
    cl::value_desc("package-name"), cl::Required, cl::cat(CommonCategory));
static cl::opt<std::string> Features(
    "features", cl::desc("features like +sve,+aes"), cl::value_desc("features"), cl::Optional, cl::cat(CommonCategory));

// JIT Options
cl::OptionCategory JITCategory("Tool JIT Options");
static cl::opt<std::string> TmplFile_(
    "tmpl", cl::desc("Tmpl file path"), cl::value_desc("tmpl-path"), cl::cat(JITCategory));

// STATIC-LINK Options
cl::OptionCategory StaticLinkCategory("Tool STATIC-LINK Options");
static cl::opt<std::string> GoProto_("goproto", cl::desc("The go file that declares go functions"),
    cl::value_desc("go-proto-path"), cl::cat(StaticLinkCategory));

bool CheckModeOptions()
{
    if (Mode_.empty()) {
        outs() << "--mode is empty\n";
        return false;
    }
    if (Mode_ != "JIT" && Mode_ != "SL") {
        outs() << "--mode is invalid\n";
        return false;
    }
    if (SourceFile_.empty()) {
        outs() << "--source is empty\n";
        return false;
    }
    if (OutputPath_.empty()) {
        outs() << "--output is empty\n";
        return false;
    }
    if (LdScript.empty()) {
        outs() << "--link-ld is empty\n";
        return false;
    }
    if (Package_.empty()) {
        outs() << "--package is empty\n";
        return false;
    }
    if (Mode_ == "JIT") {
        if (TmplFile_.empty()) {
            outs() << "tmpl is empty\n";
            return false;
        }
    }
    if (Mode_ == "SL") {
        if (GoProto_.empty()) {
            outs() << "goproto is empty\n";
            return false;
        }
    }
    return true;
}

int main(int argc, char** argv)
{
    cl::HideUnrelatedOptions({&CommonCategory, &JITCategory, &StaticLinkCategory});
    cl::ParseCommandLineOptions(argc, argv, "Tools for Go assembly conversion\n");
    if (!CheckModeOptions()) {
        return 1;
    }

    if (Debug) {
        DebugFlag = true;
    }
    std::string Mode = Mode_;
    std::string SourceFile = SourceFile_;
    std::string OutputPath = OutputPath_;
    std::string Package = Package_;
    std::string TmplFile = TmplFile_;
    std::string GoProto = GoProto_;

    auto BaseName = GetSourceName(SourceFile);
    if (BaseName.empty()) {
        return 1;
    }

    InitializeAllTargets();
    InitializeAllTargetMCs();
    InitializeAllAsmPrinters();
    InitializeAllAsmParsers();
    InitializeAllDisassemblers();

    Triple TheTriple("aarch64-linux-gnu");
    MCContextBundle Bundle(TheTriple, Features);
    FindSP(Bundle);

    // object padding 生成
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

        MCContext Ctx(TheTriple, &Bundle.getMCAsmInfo(), &Bundle.getMCRegisterInfo(), &Bundle.getMCSubtargetInfo(),
            &SrcMgr, &Bundle.getMCTargetOptions());
        std::unique_ptr<MCObjectFileInfo> MOFI(Bundle.getTarget().createMCObjectFileInfo(Ctx, false));
        Ctx.setObjectFileInfo(MOFI.get());

        auto MAB = std::unique_ptr<MCAsmBackend>(Bundle.getTarget().createMCAsmBackend(
            Bundle.getMCSubtargetInfo(), Bundle.getMCRegisterInfo(), Bundle.getMCTargetOptions()));
        auto MCE = std::unique_ptr<MCCodeEmitter>(Bundle.getTarget().createMCCodeEmitter(Bundle.getMCInstrInfo(), Ctx));

        std::error_code EC;
        raw_fd_ostream Out(ObjFile, EC, sys::fs::OF_None);
        if (EC) {
            outs() << EC.message() << "\n";
            return 1;
        }

        auto MOW = MAB->createObjectWriter(Out);
        // 不padding，会导致hasDefOfPhysReg功能失效
        auto Streamer =
            std::make_unique<PaddingNopObjectStreamer>(Ctx, std::move(MAB), std::move(MOW), std::move(MCE), Bundle);
        Streamer->initSections(false, Bundle.getMCSubtargetInfo());

        std::unique_ptr<MCAsmParser> Parser(llvm::createMCAsmParser(SrcMgr, Ctx, *Streamer, Bundle.getMCAsmInfo()));
        std::unique_ptr<MCTargetAsmParser> TAP(Bundle.getTarget().createMCAsmParser(
            Bundle.getMCSubtargetInfo(), *Parser, Bundle.getMCInstrInfo(), Bundle.getMCTargetOptions()));
        Parser->setTargetParser(*TAP);
        if (Parser->Run(false)) {
            outs() << "asm parse failed\n";
            return 1;
        }
        Streamer->finish();
    }

    // 链接elf 消除相对地址
    SmallString<256> ELFFile;
    sys::path::append(ELFFile, OutputPath, (Twine(BaseName) + ".elf").str());
    {
        // ld.lld默认会优化某些指令组合，如adrp+add在寻址相对举例小于1MB时会被优化成一条adr指令，多出来的地址变为unknown
        // 这会打断BB块的数据流，为了避免这种情况，保持链接后与输入一致，增加--no-relax
        // 或者对于JIT这种来说，代码段一般都不会大于1MB，可以直接在生成汇编时加上-mcmodel=tiny
        LdScript = "--script=" + LdScript;
        std::vector<const char*> Args = {
            "ld.lld", ObjFile.c_str(), "-o", ELFFile.c_str(), LdScript.c_str(), "--no-relax"};
        static const lld::DriverDef Drivers[] = {{lld::Gnu, lld::elf::link}};
        // 内部会调用cl::ParseCommandLineOptions，导致前面的选项被清空
        lld::Result Res = lld::lldMain(Args, outs(), errs(), Drivers);
        if (Res.retCode != 0) {
            outs() << "ld.lld failed\n";
            return 1;
        }
    }

    uint64_t DumpSize = 0;
    DumpElf(OutputPath, ELFFile, Bundle, Package, BaseName, DumpSize, Mode);

    std::string& EntryBB = BaseName;
    std::vector<BasicBlock> BBs;
    int EntryIdx = SplitBasicBlocks(Bundle, BBs, EntryBB);

    std::vector<std::vector<size_t>> CFG;
    std::vector<BBSP> BBSPVec;
    std::vector<std::pair<uint64_t, int64_t>> SPDelta;

    BuildCFG(Bundle, BBs, CFG);
    LLVM_DEBUG(if (HasCycle(BBs, CFG)) { dbgs() << "存在环\n"; });

    CalcSPDelta(Bundle, BBs, BBSPVec, SPDelta);
    auto Res = ComputeMaxSPDepth(CFG, BBSPVec);
    int i = 1;
    LLVM_DEBUG(dbgs() << "EntryIdx = " << EntryIdx << "\n";);

    auto Key = ComputeMaxSPDepthAtInst(EntryIdx, BBs, CFG, BBSPVec, SPDelta);
    LLVM_DEBUG(for (size_t i = 0; i < Key.size();
        i++) { dbgs() << "第" << i << "个SP" << SPDelta[i].first << " 最大深度:" << Key[i] << "\n"; });

    if (Mode == "JIT") {
        DumpSubr(BBs[EntryIdx], Package, OutputPath, BaseName, SPDelta, Key, DumpSize);
        DumpTmpl(TmplFile, Package, OutputPath, BaseName);
    }

    if (Mode == "SL") {
        SmallString<256> StaticFile;
        sys::path::append(StaticFile, OutputPath, (Twine(BaseName) + "_arm64.s").str());
        auto MBExp = MemoryBuffer::getFile(SourceFile);
        if (!MBExp) {
            outs() << "getFile failed\n";
            return 1;
        }
        std::unique_ptr<MemoryBuffer> MB = std::move(*MBExp);

        SourceMgr SrcMgr;
        SrcMgr.AddNewSourceBuffer(std::move(MB), SMLoc());

        MCContext Ctx(TheTriple, &Bundle.getMCAsmInfo(), &Bundle.getMCRegisterInfo(), &Bundle.getMCSubtargetInfo(),
            &SrcMgr, &Bundle.getMCTargetOptions());
        std::unique_ptr<MCObjectFileInfo> MOFI(Bundle.getTarget().createMCObjectFileInfo(Ctx, false));
        Ctx.setObjectFileInfo(MOFI.get());

        auto MAB = std::unique_ptr<MCAsmBackend>(Bundle.getTarget().createMCAsmBackend(
            Bundle.getMCSubtargetInfo(), Bundle.getMCRegisterInfo(), Bundle.getMCTargetOptions()));
        auto MCE = std::unique_ptr<MCCodeEmitter>(Bundle.getTarget().createMCCodeEmitter(Bundle.getMCInstrInfo(), Ctx));

        std::error_code EC;
        raw_fd_ostream Out(StaticFile, EC, sys::fs::OF_None);
        if (EC) {
            outs() << EC.message() << "\n";
            return 1;
        }
        raw_null_ostream NullOS;
        auto MOW = MAB->createObjectWriter(NullOS);

        auto Streamer =
            std::make_unique<Plan9Streamer>(Ctx, std::move(MAB), std::move(MOW), std::move(MCE), Out, Bundle);
        Streamer->initSections(false, Bundle.getMCSubtargetInfo());

        std::unique_ptr<MCAsmParser> Parser(llvm::createMCAsmParser(SrcMgr, Ctx, *Streamer, Bundle.getMCAsmInfo()));
        std::unique_ptr<MCTargetAsmParser> TAP(Bundle.getTarget().createMCAsmParser(
            Bundle.getMCSubtargetInfo(), *Parser, Bundle.getMCInstrInfo(), Bundle.getMCTargetOptions()));
        Parser->setTargetParser(*TAP);
        if (Parser->Run(false)) {
            outs() << "asm parse failed\n";
            return 1;
        }
        Streamer->finish();

        // 函数签名解析
        auto ParseRes = parseGoFile(GoProto);
        if (!ParseRes.error.empty()) {
            outs() << "Error: " << ParseRes.error << "\n";
            return 1;
        }

        LLVM_DEBUG(for (const auto& [name, sig] : ParseRes.funcs) {
            dbgs() << "Function: " << name << "\n";
            for (const auto& p : sig.params) {
                dbgs() << "  param: " << (p.name.empty() ? "(anon)" : p.name.c_str()) << " " << p.type << "\n";
            }
            for (const auto& r : sig.results) {
                dbgs() << "  return: " << (r.name.empty() ? "(anon)" : r.name.c_str()) << " " << r.type << "\n";
            }
            dbgs() << "\n";
        });
    }

    outs() << "ALL DONE\n";
    return 0;
}