#include "dump_elf.h"
#include "utils.h"

#include "llvm/MC/MCAsmInfo.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCDisassembler/MCDisassembler.h"
#include "llvm/MC/MCInst.h"
#include "llvm/MC/MCSubtargetInfo.h"
#include "llvm/MC/TargetRegistry.h"
#include "llvm/MC/MCInstrDesc.h"
#include "llvm/Support/Debug.h"
#include "llvm/Support/NativeFormatting.h"
#include "llvm/Support/Path.h"
#include "llvm/Support/SourceMgr.h"
#include "llvm/Support/Format.h"
#include "llvm/Support/MemoryBuffer.h"
#include "llvm/Support/LineIterator.h"
#include "llvm/Support/raw_ostream.h"
#include "llvm/Support/FileSystem.h"
#include "llvm/Object/ObjectFile.h"
#include "llvm/Object/ELFObjectFile.h"

#include <cstdint>
#include <set>
#include <sys/types.h>
#include <unordered_map>

using namespace llvm;
using namespace llvm::object;
#define DEBUG_TYPE "dump_elf"

std::vector<MCInst> Text;
std::vector<uint64_t> TextPC;
std::vector<uint32_t> TextSize;
std::unordered_map<uint64_t, size_t> Addr2Idx;
std::vector<FuncRange> Funcs;
std::set<uint64_t> Rets;

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
        LLVM_DEBUG(dbgs() << F.Name << " " << format_hex(F.StartAddr, 6) << " " << format_hex(F.EndAddr, 6) << "\n";);
    }
    llvm::sort(Funcs, [](auto &a, auto &b) { return a.StartAddr < b.StartAddr; });
}

static bool DisasmTextSection(MCContextBundle &Bundle, const ObjectFile &Obj, const SectionRef &Sec)
{
    Expected<StringRef> ContentExp = Sec.getContents();
    if (!ContentExp) {
        outs() << "getContents failed\n";
        return false;
    }

    StringRef Content = *ContentExp;
    uint64_t SectionAddr = Sec.getAddress();

    MCContext Ctx(
        Bundle.getTriple(), &Bundle.getMCAsmInfo(), &Bundle.getMCRegisterInfo(), &Bundle.getMCSubtargetInfo());
    std::unique_ptr<MCDisassembler> Disasm(Bundle.getTarget().createMCDisassembler(Bundle.getMCSubtargetInfo(), Ctx));
    if (!Disasm) {
        outs() << "create MCDisassembler failed\n";
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
            LLVM_DEBUG(PrintInstHelper(Inst, Bundle, CurAddr););
            Text.emplace_back(std::move(Inst));
            TextPC.push_back(CurAddr);
            TextSize.push_back(InstSize);
            Addr2Idx[CurAddr] = Text.size() - 1;
        }
        // 无法解析时，InstSize会存储需要跳过的字节数
        Data += InstSize;
        CurAddr += InstSize;
    }
    return true;
}

static void DumpRawBytes(std::unique_ptr<raw_ostream> &OS, const uint8_t *Data, size_t Offset, size_t Size)
{
    for (size_t i = 0; i < Size;) {
        *OS << "    ";
        size_t LineEnd = std::min(i + 16, Size);
        for (size_t j = i; j < LineEnd; ++j) {
            *OS << format_hex(Data[Offset + j], 4) << ", ";
        }
        *OS << "   // data\n";
        i = LineEnd;
    }
}

void DumpElf(const std::string &OutputPath, StringRef ElfPath, MCContextBundle &Bundle, const std::string &Package,
    const std::string &BaseName, uint64_t &DumpTextSize, const std::string &Mode)
{
    SmallString<256> DumpFile;
    sys::path::append(DumpFile, OutputPath, (Twine(BaseName) + "_text_arm64.go").str());
    auto Buf = MemoryBuffer::getFile(ElfPath);
    if (!Buf) {
        outs() << "open ELF file failed\n";
        return;
    }
    Expected<std::unique_ptr<ObjectFile>> ObjExp = ObjectFile::createObjectFile((*Buf)->getMemBufferRef());
    if (!ObjExp) {
        outs() << "createObjectFile failed\n";
        return;
    }
    ObjectFile &Obj = **ObjExp;
    CollectFuncRanges(Obj);  // 获取函数起止地址

    std::unique_ptr<raw_ostream> DumpOS;
    if (Mode == "JIT") {
        std::error_code EC;
        DumpOS = std::make_unique<raw_fd_ostream>(DumpFile, EC, sys::fs::OF_None);
        if (EC) {
            outs() << "Dump file error: " << EC.message() << "\n";
            return;
        }
    } else {
        DumpOS = std::make_unique<raw_null_ostream>();
    }

    *DumpOS << "package " << Package << "\n\n";
    *DumpOS << "var _text_" << BaseName << " = []byte{\n";

    for (auto &Sec : Obj.sections()) {
        if (!Sec.isData() && !Sec.isText() && !Sec.isBSS()) {
            continue;
        }

        Expected<StringRef> NameExp = Sec.getName();
        if (!NameExp) {
            outs() << "Get section name failed\n";
            continue;
        }
        StringRef Name = *NameExp;

        uint64_t BaseAddr = Sec.getAddress();
        uint64_t Size = Sec.getSize();

        *DumpOS << "    // " << format_hex(BaseAddr, 18) << " Contents of section " << Name << ":\n";

        if (Sec.isText()) {
            DisasmTextSection(Bundle, Obj, Sec);

            Expected<StringRef> ContentExp = Sec.getContents();
            if (!ContentExp) {
                continue;
            }
            StringRef Content = *ContentExp;
            const uint8_t *Bytes = reinterpret_cast<const uint8_t *>(Content.data());
            uint64_t BaseAddr = Sec.getAddress();
            size_t TotalSize = Content.size();
            DumpTextSize += TotalSize;

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
                *DumpOS << "    ";
                for (uint32_t j = 0; j < InstLen; ++j) {
                    *DumpOS << format_hex(Bytes[ByteIndex + j], 4) << ", ";
                }

                // 指令注释
                std::string InstStr;
                raw_string_ostream OSS(InstStr);
                Bundle.getMCInstPrinter().printInst(&Text[i], InstAddr, {}, Bundle.getMCSubtargetInfo(), OSS);
                *DumpOS << "   // " << format_hex(InstAddr, 18) << " " << InstStr << "\n";

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
            DumpTextSize += Size;
            for (uint64_t i = 0; i < Size; i += 16) {
                *DumpOS << "    ";
                uint64_t LineBytes = std::min<uint64_t>(16, Size - i);
                for (uint64_t j = 0; j < LineBytes; ++j) {
                    *DumpOS << "0x00, ";
                }
                *DumpOS << "   \n";
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
            DumpTextSize += DataSize;

            for (size_t i = 0; i < DataSize; i += 16) {
                *DumpOS << "    ";
                uint64_t LineBytes = std::min<uint64_t>(16, DataSize - i);
                for (size_t j = 0; j < LineBytes; ++j) {
                    *DumpOS << format_hex(Data[i + j], 4) << ", ";
                }
                *DumpOS << "   \n";
            }
        }
    }
    *DumpOS << "}\n";
}

void DumpSubr(const BasicBlock &EntryBB, const std::string &Package, const std::string &OutputPath,
    const std::string &BaseName, const std::vector<std::pair<uint64_t, int64_t>> &SPDelta,
    const std::vector<int64_t> &Depth, uint64_t DumpTextSize)
{
    SmallString<256> DumpFile;
    sys::path::append(DumpFile, OutputPath, (Twine(BaseName) + "_subr.go").str());
    std::error_code EC;
    raw_fd_ostream DumpOS(DumpFile, EC, sys::fs::OF_None);
    if (EC) {
        outs() << EC.message() << "\n";
        return;
    }

    DumpOS << "package " << Package << "\n\n"
           << "import (\n    `github.com/bytedance/sonic/loader`\n)\n\n"
           << "const (\n    _entry__" << BaseName << " = " << EntryBB.StartAddr << "\n)\n\n";

    int64_t MaxDepth = 0;
    for (auto x : Depth) {
        MaxDepth = std::max(MaxDepth, x);
    }
    DumpOS << "const (\n    _stack__" << BaseName << " = " << MaxDepth << "\n)\n\n"
           << "const (\n    _size__" << BaseName << " = " << DumpTextSize << "\n)\n\n"
           << "var (\n    _pcsp__" << BaseName << " = [][2]uint32{\n"
           << "        {0x1, 0},\n";

    std::map<uint64_t, int64_t> SPDump;
    for (size_t i = 1; i < SPDelta.size(); i++) {
        SPDump[SPDelta[i].first] = Depth[i];
    }
    for (auto x : Rets) {
        SPDump[x] = 0;
    }
    for (auto [Addr, SP] : SPDump) {
        DumpOS << "        {0x" << Twine::utohexstr(Addr) << ", " << SP << "},\n";
    }

    DumpOS << "        {0x" << Twine::utohexstr(DumpTextSize) << ", " << MaxDepth << "},\n"
           << "    }\n)\n\n"
           << "var _cfunc_" << BaseName << " = []loader.CFunc{\n    {\"_" << BaseName << "_entry\", 0, _entry__"
           << BaseName << ", 0, nil},\n"
           << "    {\"_" << BaseName << "\", _entry__" << BaseName << ", _size__" << BaseName << ", _stack__"
           << BaseName << ", _pcsp__" << BaseName << "},\n"
           << "}\n";
}

void DumpTmpl(
    const std::string &TmplFile, const std::string &Package, const std::string &OutputPath, const std::string &BaseName)
{
    ErrorOr<std::unique_ptr<MemoryBuffer>> BufOrErr = MemoryBuffer::getFile(TmplFile);
    if (std::error_code EC = BufOrErr.getError()) {
        outs() << "Failed to open template file '" << TmplFile << "': " + EC.message() << "\n";
        return;
    }
    MemoryBuffer &Buf = *BufOrErr.get();

    SmallString<256> OutPath;
    sys::path::append(OutPath, OutputPath, (Twine(BaseName) + ".go").str());

    std::error_code EC;
    raw_fd_ostream OutFile(OutPath, EC, sys::fs::OF_Text);
    if (EC) {
        outs() << "Failed to create output file '" << OutPath.str() << "': " << EC.message() << "\n";
    }

    bool FoundPackageLine = false;
    StringRef Placeholder = "package {{PACKAGE}}";

    for (line_iterator LineIt(Buf, false, '\0'); !LineIt.is_at_eof(); ++LineIt) {
        StringRef Line = *LineIt;
        if (!FoundPackageLine) {
            if (Line == Placeholder) {
                // 替换并输出 package 行
                OutFile << "package " << Package << '\n';
                FoundPackageLine = true;
            }
        } else {
            OutFile << Line << '\n';
        }
    }
}