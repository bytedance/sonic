#include "dump_elf.h"
#include "utils.h"

#include "llvm/MC/SectionKind.h"
#include "llvm/MC/MCAsmInfo.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCDisassembler/MCDisassembler.h"
#include "llvm/MC/MCInst.h"
#include "llvm/MC/MCInstrInfo.h"
#include "llvm/MC/MCRegisterInfo.h"
#include "llvm/MC/MCSubtargetInfo.h"
#include "llvm/MC/TargetRegistry.h"
#include "llvm/MC/MCInstrDesc.h"
#include "llvm/Support/NativeFormatting.h"
#include "llvm/Support/SourceMgr.h"
#include "llvm/Support/Format.h"
#include "llvm/Support/TargetSelect.h"
#include "llvm/Support/MemoryBuffer.h"
#include "llvm/Support/raw_ostream.h"
#include "llvm/Support/FileSystem.h"
#include "llvm/Object/ObjectFile.h"
#include "llvm/Object/ELFObjectFile.h"

#include <cstdint>
#include <sys/types.h>
#include <unordered_map>

using namespace llvm;
using namespace llvm::object;

std::string DumpFile = "lookup_small_key.dump";

std::vector<MCInst> Text;
std::vector<uint64_t> TextPC;
std::unordered_map<uint64_t, size_t> Addr2Idx;
std::vector<FuncRange> Funcs;

static void CollectFuncRanges(ObjectFile &Obj)
{
    auto *ELFObj = dyn_cast<ELFObjectFileBase>(&Obj);
    if (!ELFObj) {
        return;
    }
    for (const ELFSymbolRef &Sym : ELFObj->symbols()) {
        Expected<SymbolRef::Type> Ty = Sym.getType();
        if (!Ty || *Ty != SymbolRef::ST_Function) {
            continue;
        }

        Expected<StringRef> Name = Sym.getName();
        Expected<uint64_t> Addr = Sym.getAddress();
        Expected<uint64_t> Size = Sym.getSize();
        if (!Name || !Addr || !Size) {
            continue;
        }

        FuncRange F;
        F.StartAddr = *Addr;
        F.EndAddr = *Addr + *Size;
        F.Name = Name->str();
        Funcs.push_back(F);
        errs() << F.Name << " " << format_hex(F.StartAddr, 6) << " " << format_hex(F.EndAddr, 6) << "\n";
    }
    llvm::sort(Funcs, [](auto &a, auto &b) { return a.StartAddr < b.StartAddr; });
}

static bool DisasmTextSection(const ObjectFile &Obj, const SectionRef &Sec)
{
    Expected<StringRef> ContentExp = Sec.getContents();
    if (!ContentExp) {
        errs() << "getContents failed\n";
        return false;
    }

    StringRef Content = *ContentExp;
    uint64_t SectionAddr = Sec.getAddress();

    InitializeAllTargetInfos();
    InitializeAllTargetMCs();
    InitializeAllDisassemblers();

    Triple TheTriple = Obj.makeTriple();
    StringRef TripleName = TheTriple.getTriple();
    std::string Error;
    const Target *TheTarget = TargetRegistry::lookupTarget(TripleName, Error);
    if (!TheTarget) {
        errs() << "lookupTarget failed: " << Error << "\n";
        return false;
    }

    std::unique_ptr<MCRegisterInfo> MRI(TheTarget->createMCRegInfo(TheTriple.str()));
    assert(MRI && "Unable to create MCRegisterInfo!");
    MCTargetOptions MCOptions;
    std::unique_ptr<MCAsmInfo> MAI(TheTarget->createMCAsmInfo(*MRI, TheTriple.str(), MCOptions));
    assert(MAI && "Unable to create MCAsmInfo!");
    std::unique_ptr<MCSubtargetInfo> STI(TheTarget->createMCSubtargetInfo(TheTriple.str(), "generic", "+sve"));
    assert(STI && "Unable to create MCSubtargetInfo!");
    std::unique_ptr<MCInstrInfo> MCII(TheTarget->createMCInstrInfo());
    assert(MCII && "Unable to create MCInstrInfo!");

    MCContext Ctx(TheTriple, MAI.get(), MRI.get(), STI.get());
    std::unique_ptr<MCDisassembler> Disasm(TheTarget->createMCDisassembler(*STI, Ctx));
    if (!Disasm) {
        errs() << "create MCDisassembler failed\n";
        return false;
    }

    uint64_t CurAddr = SectionAddr;
    const uint8_t *Data = reinterpret_cast<const uint8_t *>(Content.data());
    const uint8_t *End = Data + Content.size();
    while (Data < End) {
        MCInst Inst;
        uint64_t InstSize = 0;
        // 每次反汇编都需要传入剩余全部字节
        ArrayRef<uint8_t> Bytes(Data, End - Data);

        auto DisasmStat = Disasm->getInstruction(Inst, InstSize, Bytes, CurAddr, errs());
        if (DisasmStat == llvm::MCDisassembler::DecodeStatus::Success) {
            PrintInstHelper(Inst, MRI.get(), MCII.get(), CurAddr);
            Text.emplace_back(std::move(Inst));
            TextPC.push_back(CurAddr);
            Addr2Idx[CurAddr] = Text.size() - 1;
        }
        Data += InstSize;
        CurAddr += InstSize;
    }
    return true;
}

void DumpElf(llvm::StringRef ElfPath)
{
    auto Buf = MemoryBuffer::getFile(ElfPath);
    if (!Buf) {
        errs() << "open ELF file failed\n";
        return;
    }
    Expected<std::unique_ptr<ObjectFile>> ObjExp = ObjectFile::createObjectFile((*Buf)->getMemBufferRef());
    if (!ObjExp) {
        errs() << "createObjectFile failed\n";
        return;
    }
    ObjectFile &Obj = **ObjExp;
    CollectFuncRanges(Obj);  // 获取函数起止地址

    std::error_code EC;
    raw_fd_ostream DumpOS(DumpFile, EC, sys::fs::OF_None);
    if (EC) {
        errs() << EC.message() << "\n";
        return;
    }

    for (auto &Sec : Obj.sections()) {
        if (!Sec.isData() && !Sec.isText()) {
            continue;
        }

        Expected<StringRef> NameExp = Sec.getName();
        if (!NameExp) {
            errs() << "Get section name failed\n";
            continue;
        }

        StringRef Name = *NameExp;
        if (!Name.starts_with(".text") && !Name.starts_with(".rodata")) {
            continue;
        }

        if (Name.starts_with(".text")) {
            DisasmTextSection(Obj, Sec);
        }
        DumpOS << "Contents of section " << Name << ":\n";

        Expected<StringRef> ContentExp = Sec.getContents();
        if (!ContentExp) {
            errs() << "getContents failed\n";
            continue;
        }
        StringRef Content = *ContentExp;
        uint64_t BaseAddr = Sec.getAddress();
        size_t Size = Content.size();
        size_t InstCount = Size / 4;
        size_t RemainInst = (4 - (InstCount % 4)) % 4;
        for (size_t i = 0; i < Size; i += 16) {
            DumpOS << format("%08" PRIx64 " ", BaseAddr + i);
            for (size_t j = 0; j < 16; j++) {
                if (i + j < Size) {
                    DumpOS << format(" %02x", (uint8_t)Content[i + j]);
                } else {
                    DumpOS << "   ";
                }
            }
            DumpOS << "   ";
            for (size_t j = 0; j < 16 && i + j < Size; j++) {
                char CH = Content[i + j];
                DumpOS << (isprint(CH) ? CH : '.');
            }
            DumpOS << "\n";
        }
    }
}