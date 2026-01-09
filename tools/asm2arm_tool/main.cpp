#include "mc_bundle.h"
#include "streamer.h"
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
#include "llvm/Support/MemoryBuffer.h"
#include "llvm/Support/Program.h"
#include "llvm/Support/SourceMgr.h"
#include "llvm/Support/TargetSelect.h"
#include "llvm/Support/VirtualFileSystem.h"
#include "llvm/Support/raw_ostream.h"
#include "llvm/TargetParser/Host.h"
#include "llvm/TargetParser/Triple.h"

#include <memory>
#include <optional>
#include <string>
#include <vector>

using namespace llvm;

static cl::opt<std::string> SourceFile(
    "source-file", cl::desc("input .cxx file"), cl::value_desc("source-file-path"), cl::Required);
static cl::opt<std::string> LdScript(
    "link-ld", cl::desc("linker script"), cl::value_desc("link-ld-path"), cl::Required);
static cl::opt<std::string> Package("package", cl::desc("The package to which the generated Go file belongs"),
    cl::value_desc("package-name"), cl::Required);
static cl::opt<std::string> TmplDir(
    "TmplDir", cl::desc("Folder where Tmpl files are stored"), cl::value_desc("Tmpl-files-Dir"), cl::Required);

int main(int argc, char **argv)
{
    cl::ParseCommandLineOptions(argc, argv, "assembly->object->elf->objdump\n");

    auto BaseName = GetSourceName(SourceFile);
    if (BaseName.empty()) {
        return 1;
    }

    InitializeAllTargets();
    InitializeAllTargetMCs();
    InitializeAllAsmPrinters();
    InitializeAllAsmParsers();
    InitializeAllDisassemblers();

    // 汇编生成
    std::string AsmFile = BaseName + ".s";
    {
        std::vector<StringRef> Args = {
            "clang", "-S", SourceFile, "-o", AsmFile, "-O2", "-D__SVE__", "-march=armv8-a+sve", "-I/usr/include/simde"};
        std::string Err;
        int RC = sys::ExecuteAndWait("/home/yupan/.local/LLVM19/bin/clang", Args, std::nullopt, {}, 0, 0, &Err);
        if (RC) {
            errs() << "clang failed: " << Err << "\n";
            return 1;
        }
    }

    Triple TheTriple("aarch64-linux-gnu");
    MCContextBundle Bundle(TheTriple);

    // object padding 生成
    std::string ObjFile = BaseName + ".o";
    {
        auto MBExp = MemoryBuffer::getFile(AsmFile);
        if (!MBExp) {
            errs() << "getFile failed\n";
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
            errs() << EC.message() << "\n";
            return 1;
        }

        auto MOW = MAB->createObjectWriter(Out);
        auto Streamer =
            std::make_unique<PaddingNopObjectStreamer>(Ctx, std::move(MAB), std::move(MOW), std::move(MCE), Bundle);
        Streamer->initSections(false, Bundle.getMCSubtargetInfo());

        std::unique_ptr<MCAsmParser> Parser(llvm::createMCAsmParser(SrcMgr, Ctx, *Streamer, Bundle.getMCAsmInfo()));
        std::unique_ptr<MCTargetAsmParser> TAP(Bundle.getTarget().createMCAsmParser(
            Bundle.getMCSubtargetInfo(), *Parser, Bundle.getMCInstrInfo(), Bundle.getMCTargetOptions()));
        Parser->setTargetParser(*TAP);
        if (Parser->Run(false)) {
            errs() << "asm parse failed\n";
            return 1;
        }
        Streamer->finish();
    }

    // 链接elf 消除相对地址
    std::string ELFFile = BaseName + ".elf";
    {
        std::vector<StringRef> Args = {"ld.lld", ObjFile, "-o", ELFFile, "-T", LdScript};
        std::string Err;
        int RC = sys::ExecuteAndWait("/home/yupan/.local/LLVM19/bin/ld.lld", Args, std::nullopt, {}, 0, 0, &Err);
        if (RC) {
            errs() << "ld.lld failed: " << Err << "\n";
            return 1;
        }
    }

    uint64_t TextStartAddr;
    uint64_t DumpSize;
    DumpElf(ELFFile, Bundle, Package, BaseName, TextStartAddr, DumpSize);

    std::string &EntryBB = BaseName;
    std::vector<BasicBlock> BBs;
    int EntryIdx = SplitBasicBlocks(Bundle, BBs, EntryBB);

    std::vector<std::vector<size_t>> CFG;
    std::vector<BBSP> BBSPVec;
    std::vector<std::pair<uint64_t, int64_t>> SPDelta;

    BuildCFG(Bundle, BBs, CFG);
    if (HasCycle(BBs, CFG)) {
        outs() << "存在环\n";
    }

    CalcSPDelta(Bundle, BBs, BBSPVec, SPDelta);
    auto Res = ComputeMaxSPDepth(CFG, BBSPVec);
    int i = 1;
    for (auto x : Res) {
        outs() << "第" << i++ << "块为入口时最大SP: " << x << "\n";
    }
    outs() << "EntryIdx = " << EntryIdx << "\n";

    auto Key = ComputeMaxSPDepthAtInst(EntryIdx, BBs, CFG, BBSPVec, SPDelta);
    for (size_t i = 0; i < Key.size(); i++) {
        outs() << "第" << i << "个SP" << SPDelta[i].first << " 最大深度:" << Key[i] << "\n";
    }

    DumpSubr(BBs[EntryIdx], Package, BaseName, SPDelta, Key, TextStartAddr, DumpSize);
    DumpTmpl(TmplDir, Package, BaseName);

    // llvm-objdump
    std::string TextFile = BaseName + ".text";
    {
        std::vector<StringRef> Args = {"llvm-objdump", "-S", ELFFile};
        std::error_code EC;
        raw_fd_ostream TextOS(TextFile, EC, sys::fs::OF_None);
        if (EC) {
            errs() << EC.message() << "\n";
            return 1;
        }

        std::string Err;
        int RC = sys::ExecuteAndWait("/home/yupan/.local/LLVM19/bin/llvm-objdump", Args, std::nullopt,
            {std::nullopt, TextFile, std::nullopt}, 0, 0, &Err, nullptr);
        if (RC) {
            errs() << "llvm-objdump failed: " << Err << "\n";
            return 1;
        }
    }

    outs() << "ALL DONE\n";
    return 0;
}