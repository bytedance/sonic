#include "mc_bundle.h"
#include "streamer.h"
#include "plan9_streamer.h"
#include "dump_elf.h"
#include "cal_depth.h"
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

#include <memory>
#include <string>
#include <vector>

using namespace llvm;
#define DEBUG_TYPE "main"
LLD_HAS_DRIVER(elf)

cl::OptionCategory JITCategory("JIT Options");

static cl::opt<bool> Debug("debug", cl::desc("Enable debug output"), cl::init(false), cl::cat(JITCategory));
static cl::opt<std::string> SourceFile_(
    "source", cl::desc("input ASM file"), cl::value_desc("ASM-file-path"), cl::Required, cl::cat(JITCategory));
static cl::opt<std::string> OutputPath_(
    "output", cl::desc("Output path of *.go files"), cl::value_desc("output-path"), cl::Required, cl::cat(JITCategory));
static cl::opt<std::string> LdScript(
    "link-ld", cl::desc("linker script"), cl::value_desc("link-ld-path"), cl::Required, cl::cat(JITCategory));
static cl::opt<std::string> Package_("package", cl::desc("The package to which the generated Go file belongs"),
    cl::value_desc("package-name"), cl::Required, cl::cat(JITCategory));
static cl::opt<std::string> TmplDir_("TmplDir", cl::desc("Folder where Tmpl files are stored"),
    cl::value_desc("Tmpl-files-Dir"), cl::Required, cl::cat(JITCategory));
static cl::opt<std::string> Features(
    "features", cl::desc("features like +sve,+aes"), cl::value_desc("features"), cl::Optional, cl::cat(JITCategory));

int main(int argc, char **argv)
{
    cl::HideUnrelatedOptions(JITCategory);
    cl::ParseCommandLineOptions(argc, argv, "assembly->object->elf->objdump\n");
    if (Debug) {
        DebugFlag = true;
    }
    std::string SourceFile = SourceFile_;
    std::string OutputPath = OutputPath_;
    std::string Package = Package_;
    std::string TmplDir = TmplDir_;

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
    SmallString<256> StaticFile;
    sys::path::append(StaticFile, OutputPath, (Twine(BaseName) + "_arm64.s").str());
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
    }

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
        std::vector<const char *> Args = {
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
    DumpElf(OutputPath, ELFFile, Bundle, Package, BaseName, DumpSize);

    std::string &EntryBB = BaseName;
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

    DumpSubr(BBs[EntryIdx], Package, OutputPath, BaseName, SPDelta, Key, DumpSize);
    DumpTmpl(TmplDir, Package, OutputPath, BaseName);

    outs() << "ALL DONE\n";
    return 0;
}