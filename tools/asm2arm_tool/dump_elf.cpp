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

std::vector<MCInst> Text;
std::vector<uint64_t> TextPC;
std::vector<uint32_t> TextSize;
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
            TextSize.push_back(InstSize);
            Addr2Idx[CurAddr] = Text.size() - 1;
            Data += InstSize;
            CurAddr += InstSize;
        } else {
            Data += 1;
            CurAddr += 1;
        }
    }
    return true;
}

static void DumpRawBytes(raw_fd_ostream &OS, const uint8_t *Data, size_t Offset, size_t Size)
{
    for (size_t i = 0; i < Size;) {
        OS << "    ";
        size_t LineEnd = std::min(i + 16, Size);
        for (size_t j = i; j < LineEnd; ++j) {
            OS << format_hex(Data[Offset + j], 4) << ", ";
        }
        OS << "   // data\n";
        i = LineEnd;
    }
}

void DumpElf(StringRef ElfPath, StringRef DumpFile, const llvm::MCSubtargetInfo *STI, MCInstPrinter *MCIP,
    const std::string &Package, const std::string &BaseName)
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

    DumpOS << "package " << Package << "\n\n";
    DumpOS << "var _text_" << BaseName << " = []byte{\n";

    for (auto &Sec : Obj.sections()) {
        if (!Sec.isData() && !Sec.isText() && !Sec.isBSS()) {
            continue;
        }

        Expected<StringRef> NameExp = Sec.getName();
        if (!NameExp) {
            errs() << "Get section name failed\n";
            continue;
        }
        StringRef Name = *NameExp;

        uint64_t BaseAddr = Sec.getAddress();
        uint64_t Size = Sec.getSize();

        DumpOS << "    // " << format_hex(BaseAddr, 18) << " Contents of section " << Name << ":\n";

        if (Sec.isText()) {
            DisasmTextSection(Obj, Sec);

            Expected<StringRef> ContentExp = Sec.getContents();
            if (!ContentExp) {
                continue;
            }
            StringRef Content = *ContentExp;
            const uint8_t *Bytes = reinterpret_cast<const uint8_t *>(Content.data());
            uint64_t BaseAddr = Sec.getAddress();
            size_t TotalSize = Content.size();

            size_t NumInsts = Text.size();

            // 开头
            uint64_t CurrentAddr = BaseAddr;
            size_t ByteIndex = 0;

            if (NumInsts > 0 && TextPC[0] > BaseAddr) {
                size_t GapSize = TextPC[0] - BaseAddr;
                DumpRawBytes(DumpOS, Bytes, ByteIndex, GapSize);
                ByteIndex += GapSize;
                CurrentAddr += GapSize;
            }

            for (size_t i = 0; i < NumInsts; ++i) {
                uint64_t InstAddr = TextPC[i];
                uint32_t InstLen = TextSize[i];

                assert(InstAddr == CurrentAddr && "Address misalignment!");
                assert(ByteIndex + InstLen <= TotalSize && "Instruction overflows section");

                // 输出指令字节（单行）
                DumpOS << "    ";
                for (uint32_t j = 0; j < InstLen; ++j) {
                    DumpOS << format_hex(Bytes[ByteIndex + j], 4) << ", ";
                }

                // 指令注释
                std::string InstStr;
                raw_string_ostream OSS(InstStr);
                MCIP->printInst(&Text[i], InstAddr, {}, *STI, OSS);
                DumpOS << "   // " << InstStr << "\n";

                ByteIndex += InstLen;
                CurrentAddr += InstLen;

                // 计算到下一条指令（或段尾）的 gap
                uint64_t NextInstAddr = (i + 1 < NumInsts) ? TextPC[i + 1] : (BaseAddr + TotalSize);
                if (CurrentAddr < NextInstAddr) {
                    size_t GapSize = NextInstAddr - CurrentAddr;
                    DumpRawBytes(DumpOS, Bytes, ByteIndex, GapSize);
                    ByteIndex += GapSize;
                    CurrentAddr = NextInstAddr;
                }
            }
        } else if (Sec.isBSS()) {
            // .bss: 全零
            for (uint64_t i = 0; i < Size; i += 16) {
                DumpOS << "    ";
                uint64_t LineBytes = std::min<uint64_t>(16, Size - i);
                for (uint64_t j = 0; j < LineBytes; ++j) {
                    DumpOS << "0x00, ";
                }
                DumpOS << "   \n";
            }
        } else {
            // .data / .rodata
            Expected<StringRef> ContentExp = Sec.getContents();
            if (!ContentExp) {
                continue;
            }
            StringRef Content = *ContentExp;
            const uint8_t *Data = reinterpret_cast<const uint8_t *>(Content.data());
            size_t DataSize = Content.size();

            for (size_t i = 0; i < DataSize; i += 16) {
                DumpOS << "    ";
                uint64_t LineBytes = std::min<uint64_t>(16, DataSize - i);
                for (size_t j = 0; j < LineBytes; ++j) {
                    DumpOS << format_hex(Data[i + j], 4) << ", ";
                }
                DumpOS << "   \n";
            }
        }
    }
    DumpOS << "}\n";
}