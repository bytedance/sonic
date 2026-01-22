#include "plan9_streamer.h"
#include "utils.h"

#include "llvm/ADT/SmallVector.h"
#include "llvm/ADT/StringExtras.h"
#include "llvm/MC/MCAssembler.h"
#include "llvm/MC/MCELFStreamer.h"
#include "llvm/MC/MCFixup.h"
#include "llvm/MC/MCInst.h"
#include "llvm/Support/Debug.h"
#include "llvm/Support/Format.h"
#include "llvm/Support/raw_ostream.h"
#include <cstddef>
#include <string>
#include <unordered_map>

using namespace llvm;
#define DEBUG_TYPE "plan9_streamer"

Plan9Streamer::Plan9Streamer(llvm::MCContext &Context, std::unique_ptr<llvm::MCAsmBackend> TAB,
    std::unique_ptr<llvm::MCObjectWriter> OW, std::unique_ptr<llvm::MCCodeEmitter> Emitter, llvm::raw_fd_ostream &Out,
    MCContextBundle &Bundle)
    : MCELFStreamer(Context, std::move(TAB), std::move(OW), std::move(Emitter)), Out(Out), Bundle(Bundle)
{}

void Plan9Streamer::finish()
{
    this->flushPendingBytes();
    MCELFStreamer::finish();
}

void Plan9Streamer::emitLabel(MCSymbol *Sym, SMLoc Loc)
{
    this->flushPendingBytes();
    if (IsTopEmit == 0) {
        OutLabel(this->Out, Sym->getName()) << ":\n";

        LLVM_DEBUG({
            dbgs() << "LABEL: ";
            if (Sym->isInSection()) {
                dbgs() << Sym->getSection().getName() << "::";
            }
            dbgs() << Sym->getName() << "\n";
        });
    }
    IsTopEmit++;
    MCELFStreamer::emitLabel(Sym, Loc);
    IsTopEmit--;
}

uint32_t readLittleEndianU32(const SmallVectorImpl<char> &CB)
{
    assert(CB.size() == 4);
    return support::endian::read<uint32_t>(reinterpret_cast<const uint8_t *>(CB.data()), llvm::endianness::little);
}

std::string ToPlan9Reg(const std::string &armReg)
{
    // 去除前导非数字字符，提取数字部分
    std::string numPart;
    for (char c : armReg) {
        if (std::isdigit(c)) {
            numPart += c;
        }
    }
    return "R" + numPart;
}

std::string ToUpper(const std::string &Str)
{
    std::string Res(Str.length(), 0);
    for (size_t i = 0; i < Str.length(); i++) {
        Res[i] = toUpper(Str[i]);
    }
    return Res;
};

static std::unordered_map<std::string, std::string> BranchMap = {
    {"b", "B"},
    {"bl", "BL"},
    // 根据寄存器内的地址跳转，感觉单文件的功能不会出现
    {"blr", "BLR"},
    {"br", "BR"},

    {"b.eq", "BEQ"},
    {"b.ne", "BNE"},
    {"b.cs", "BCS"},
    {"b.hs", "BHS"},  // cs 同 hs
    {"b.cc", "BCC"},
    {"b.lo", "BLO"},  // cc 同 lo
    {"b.mi", "BMI"},
    {"b.pl", "BPL"},
    {"b.vs", "BVS"},
    {"b.vc", "BVC"},
    {"b.hi", "BHI"},
    {"b.ls", "BLS"},
    {"b.ge", "BGE"},
    {"b.lt", "BLT"},
    {"b.gt", "BGT"},
    {"b.le", "BLE"},
    {"b.al", "B"},  // 总是真
    {"b.nv", "B"},  // b.nv 应该是不存在的

    {"bc.eq", "BEQ"},
    {"bc.ne", "BNE"},
    {"bc.cs", "BCS"},
    {"bc.hs", "BHS"},
    {"bc.cc", "BCC"},
    {"bc.lo", "BLO"},
    {"bc.mi", "BMI"},
    {"bc.pl", "BPL"},
    {"bc.vs", "BVS"},
    {"bc.vc", "BVC"},
    {"bc.hi", "BHI"},
    {"bc.ls", "BLS"},
    {"bc.ge", "BGE"},
    {"bc.lt", "BLT"},
    {"bc.gt", "BGT"},
    {"bc.le", "BLE"},
    {"bc.al", "B"},
    {"bc.nv", "B"},
};

/// cbz --> cmp + beq | cbnz --> cmp + bne
/// tbz --> tst + beq | tbnz --> tst + bne
bool Plan9Streamer::makeCmpareBranch(const std::vector<std::string> &Token, const std::string &InstStr)
{
    auto &Op = Token[0];
    if (Op != "cbz" && Op != "cbnz" && Op != "tbz" && Op != "tbnz") {
        return false;
    }
    if (Op == "cbz") {
        this->Out << "    CMP $0, R" << Token[1].substr(1) << "\n";
        this->Out << "    BEQ ";
        OutLabel(this->Out, Token[2]) << "  // " << InstStr << "\n";
    } else if (Op == "cbnz") {
        this->Out << "    CMP $0, R" << Token[1].substr(1) << "\n";
        this->Out << "    BNE ";
        OutLabel(this->Out, Token[2]) << "  // " << InstStr << "\n";
    } else if (Op == "tbz") {
        this->Out << "    TST $(1<<" << Token[2].substr(1) << "), R" << Token[1].substr(1) << "\n";
        this->Out << "    BEQ ";
        OutLabel(this->Out, Token[3]) << "  // " << InstStr << "\n";
    } else {
        this->Out << "    TST $(1<<" << Token[2].substr(1) << "), R" << Token[1].substr(1) << "\n";
        this->Out << "    BNE ";
        OutLabel(this->Out, Token[3]) << "  // " << InstStr << "\n";
    }
    return true;
}

/// 跳转指令
void Plan9Streamer::makeBranch(const std::vector<std::string> &Token, const std::string &InstStr)
{
    auto &Op = Token[0];
    this->Out << "    " << BranchMap[Op] << " ";
    auto &Label = Token[1];
    OutLabel(this->Out, Label);
    this->Out << "  // " << InstStr << "\n";
}

void Plan9Streamer::makeBranchInst(const std::vector<std::string> &Token, const std::string &InstStr)
{
    auto &Op = Token[0];
    if (BranchMap.find(Op) != BranchMap.end()) {
        this->makeBranch(Token, InstStr);
        return;
    }
    if (this->makeCmpareBranch(Token, InstStr)) {
        return;
    }
    outs() << "Unsupported Branch Instruction\n";
}

void Plan9Streamer::emitInstruction(const MCInst &Inst, const MCSubtargetInfo &STI)
{
    if (IsTopEmit == 0) {
        const auto &Desc = Bundle.getMCInstrInfo().get(Inst.getOpcode());
        std::string InstStr;
        raw_string_ostream OS(InstStr);
        Bundle.getMCInstPrinter().printInst(&Inst, 0, "", STI, OS);

        SmallVector<char> Buffer;
        SmallVector<MCFixup> Fixup;
        MCELFStreamer::getAssembler().getEmitter().encodeInstruction(Inst, Buffer, Fixup, Bundle.getMCSubtargetInfo());
        auto Token = TokenizeInstruction(InstStr);
        // Fixup非空时，说明指令中存在需要在链接时处理的label参数
        // label参数在MCOperand中的判断是isExpr()，暂不清楚这种指令能否直接使用WORD表示
        if (Desc.isBranch()) {
            this->makeBranchInst(Token, InstStr);
        } else if (Token[0] == "adrp") {
            this->Out << "    ADR ";
            OutLabel(this->Out, Token[2]) << ", R" << Token[1].substr(1) << "\n";
        } else {
            this->Out << "    WORD $" << format_hex(readLittleEndianU32(Buffer), 10) << "  // " << InstStr << "\n";
        }

        LLVM_DEBUG({
            dbgs() << "INSTRUCTION: " << InstStr << " Size:" << Buffer.size() << " Fixup:" << Fixup.empty() << "\n";
        });
    }
    IsTopEmit++;
    MCELFStreamer::emitInstruction(Inst, STI);
    IsTopEmit--;
}

void Plan9Streamer::makeWordData(uint64_t Value, unsigned Size, unsigned Repeat)
{
    // 将 Value 按 Size 字节小端写入缓冲区
    auto appendBytes = [&](uint64_t Val, unsigned S) {
        for (unsigned i = 0; i < S; ++i) {
            this->WordData += static_cast<char>(Val & 0xFF);
            Val >>= 8;
        }
    };

    // 写入 Repeat 次
    for (unsigned r = 0; r < Repeat; ++r) {
        appendBytes(Value, Size);
    }

    while (this->WordData.size() >= 4) {
        // 小端WORD
        uint32_t Word =
            (static_cast<uint8_t>(this->WordData[0]) << 0) | (static_cast<uint8_t>(this->WordData[1]) << 8) |
            (static_cast<uint8_t>(this->WordData[2]) << 16) | (static_cast<uint8_t>(this->WordData[3]) << 24);

        Out << "    WORD $" << format_hex(Word, 10) << "\n";
        // 移除已输出的 4 字节
        this->WordData.erase(0, 4);
    }
}

void Plan9Streamer::flushPendingBytes()
{
    if (this->WordData.empty()) {
        return;
    }
    // 补零到 4 字节对齐
    while (this->WordData.size() % 4 != 0) {
        this->WordData += '\0';
    }
    // 输出所有完整 WORD
    while (this->WordData.size() >= 4) {
        uint32_t Word =
            (static_cast<uint8_t>(this->WordData[0]) << 0) | (static_cast<uint8_t>(this->WordData[1]) << 8) |
            (static_cast<uint8_t>(this->WordData[2]) << 16) | (static_cast<uint8_t>(this->WordData[3]) << 24);
        this->Out << "    WORD $" << format_hex(Word, 10) << "\n";
        WordData.erase(0, 4);
    }
}

void Plan9Streamer::emitIntValue(uint64_t Value, unsigned Size)
{
    if (IsTopEmit == 0) {
        this->makeWordData(Value, Size);
        LLVM_DEBUG(dbgs() << "INT DATA: .int" << (Size * 8) << " = " << Value << "\n");
    }
    IsTopEmit++;
    MCELFStreamer::emitIntValue(Value, Size);
    IsTopEmit--;
}

void Plan9Streamer::emitFill(const llvm::MCExpr &NumBytes, uint64_t FillValue, llvm::SMLoc Loc)
{
    if (IsTopEmit == 0) {
        int64_t NumBytesVal;
        bool Evaluated = NumBytes.evaluateAsAbsolute(NumBytesVal);
        this->makeWordData(FillValue, 1, NumBytesVal);
        LLVM_DEBUG({
            dbgs() << "FILL (form 1): ";
            if (Evaluated) {
                dbgs() << ".space / .zero " << NumBytesVal << " bytes, fill value = " << FillValue;
            } else {
                dbgs() << ".space <expr>, fill value = " << FillValue;
            }
            dbgs() << "\n";
        });
    }
    IsTopEmit++;
    MCELFStreamer::emitFill(NumBytes, FillValue, Loc);
    IsTopEmit--;
}

void Plan9Streamer::emitBytes(StringRef Data)
{
    if (IsTopEmit == 0) {
        this->WordData.append(Data.begin(), Data.end());
        while (this->WordData.size() >= 4) {
            uint32_t Word =
                (static_cast<uint8_t>(this->WordData[0]) << 0) | (static_cast<uint8_t>(this->WordData[1]) << 8) |
                (static_cast<uint8_t>(this->WordData[2]) << 16) | (static_cast<uint8_t>(this->WordData[3]) << 24);
            this->Out << "    WORD $" << format_hex(Word, 10) << "\n";
            this->WordData.erase(0, 4);
        }
        LLVM_DEBUG(dbgs() << "BYTE DATA: len=" << Data.size() << "\n");
    }
    IsTopEmit++;
    MCELFStreamer::emitBytes(Data);
    IsTopEmit--;
}

void Plan9Streamer::emitIdent(llvm::StringRef IdentString)
{
    LLVM_DEBUG({ dbgs() << "IDENT: .ident \"" << IdentString << "\"\n"; });
    IsTopEmit++;
    MCELFStreamer::emitIdent(IdentString);
    IsTopEmit--;
}

void Plan9Streamer::emitValueToAlignment(
    llvm::Align Alignment, int64_t Value, unsigned ValueSize, unsigned MaxBytesToEmit)
{
    LLVM_DEBUG(dbgs() << "Value Align: Alignment=" << Alignment.value() << " Value=" << Value
                      << " ValueSize=" << ValueSize << " MaxBytesToEmit=" << MaxBytesToEmit << "\n");
    IsTopEmit++;
    MCELFStreamer::emitValueToAlignment(Alignment, Value, ValueSize, MaxBytesToEmit);
    IsTopEmit--;
}

void DumpDeclareHead(llvm::raw_fd_ostream &Out, const std::string &BaseName, int64_t MaxDepth)
{
    Out << "#include \"go_asm.h\"\n"
        << "#include \"funcdata.h\"\n"
        << "#include \"textflag.h\"\n";

    int64_t GoDepth = MaxDepth < 16 ? 0 : MaxDepth - 16;
    Out << "\nTEXT ·__" << BaseName << "_entry__(SB), NOSPLIT, $" << GoDepth << "\n"
        << "    NO_LOCAL_POINTERS\n"
        << "    WORD $0x100000a0 // adr x0, .+20\n"
        << "    MOVD R0, ret(FP)\n"
        << "    RET\n\n";
}

void DumpDeclareTail(llvm::raw_fd_ostream &Out, const std::string &BaseName, ParseResult &ParseRes, int64_t MaxDepth)
{
    const auto &func = ParseRes.funcs["__" + BaseName];
    if (!func.isAllocated()) {
        outs() << BaseName << " isAllocated(): false\n";
        return;
    }
    Out << "\nTEXT ·__" << BaseName << "(SB), NOSPLIT, $0-" << func.argspace() << "\n"
        << "    NO_LOCAL_POINTERS\n";

    int64_t CheckDepth = MaxDepth + 64;

    if (CheckDepth != 0) {
        Out << "\n_entry:\n"
            << "    MOVD 16(g), R16\n";
        if (MaxDepth > 0) {
            if (MaxDepth < ((1 << 12) - 1)) {
                Out << "    SUB $" << CheckDepth << ", RSP, R17\n";
            } else if (MaxDepth < ((1 << 16) - 1)) {
                Out << "    MOVD $" << CheckDepth << ", R17\n"
                    << "    SUB R17, RSP, R17\n";
            } else {
                outs() << "too large stack size: " << CheckDepth << "\n";
                return;
            }
            Out << "    CMP R16, R17\n";
        } else {
            Out << "    CMP R16, RSP\n";
        }
        Out << "    BLS _stack_grow\n";
    }

    Out << "\n    _" << BaseName << ":\n";
    size_t offset = 0;
    for (auto &p : func.params) {
        if (p.creg.name[0] == 'x') {
            Out << "    MOVD " << p.name << "+" << offset << "(FP), R" << p.creg.name.substr(1) << "\n";
        }
        if (p.creg.name[0] == 'd') {
            Out << "    FMOVD " << p.name << "+" << offset << "(FP), F" << p.creg.name.substr(1) << "\n";
        }
        offset += p.size;
    }
    Out << "    MOVD ·_subr__" << BaseName << "(SB), R11\n"
        << "    BL (R11)\n";

    if (!func.results.empty()) {
        auto &p = func.results[0];
        if (p.creg.name[0] == 'x') {
            Out << "    MOVD R0, " << p.name << "+" << offset << "(FP)\n";
        }
        if (p.creg.name[0] == 'd') {
            Out << "    FMOVD F0, " << p.name << "+" << offset << "(FP)\n";
        }
    }
    Out << "    RET\n";

    if (CheckDepth != 0) {
        Out << "\n_stack_grow:\n"
            << "    MOVD R30, R3\n"
            << "    CALL runtime·morestack_noctxt<>(SB)\n"
            << "    JMP  _entry\n";
    }
}