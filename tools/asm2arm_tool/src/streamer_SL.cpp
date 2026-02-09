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

#include "streamer_SL.h"
#include "utils.h"

#include "llvm/ADT/SmallVector.h"
#include "llvm/ADT/StringExtras.h"
#include "llvm/MC/MCAssembler.h"
#include "llvm/MC/MCELFStreamer.h"
#include "llvm/MC/MCFixup.h"
#include "llvm/MC/MCInst.h"
#include "llvm/Support/Debug.h"
#include "llvm/Support/FileSystem.h"
#include "llvm/Support/Format.h"
#include "llvm/Support/Path.h"
#include "llvm/Support/raw_ostream.h"
#include <cstddef>
#include <string>
#include <unordered_map>

using namespace llvm;
#define DEBUG_TYPE "plan9_streamer"

namespace tool {
namespace asm2arm {

SLStreamer::SLStreamer(llvm::MCContext &Context,
                       std::unique_ptr<llvm::MCAsmBackend> AsmBackend,
                       std::unique_ptr<llvm::MCObjectWriter> ObjectWriter,
                       std::unique_ptr<llvm::MCCodeEmitter> Emitter,
                       llvm::raw_fd_ostream &Out,
                       tool::mc::MCContextBundle &Bundle,
                       const std::string &BaseName)
    : MCELFStreamer(Context, std::move(AsmBackend), std::move(ObjectWriter),
                    std::move(Emitter)),
      Out(Out), Bundle(Bundle), BaseName(BaseName) {}

uint64_t SLStreamer::GetStartProgramCounter() {
  return this->StartProgramCounter;
}

void SLStreamer::finish() {
  this->FlushPendingBytes();
  MCELFStreamer::finish();
}

void SLStreamer::emitLabel(MCSymbol *Sym, SMLoc Loc) {
  this->FlushPendingBytes();
  if (IsTopEmit == 0) {
    if (!Sym->getName().empty()) {
      tool::OutLabel(this->Out, Sym->getName()) << ":\n";
    }

    LLVM_DEBUG({
      dbgs() << "LABEL: ";
      if (Sym->isInSection()) {
        dbgs() << Sym->getSection().getName() << "::";
      }
      dbgs() << Sym->getName() << "\n";
    });
  }
  if (Sym->getName() == BaseName) {
    this->StartProgramCounter = this->ProgramCounter;
  }
  IsTopEmit++;
  MCELFStreamer::emitLabel(Sym, Loc);
  IsTopEmit--;
}

/**
 * @brief 读取小端序的32位无符号整数
 *
 * @param CB 字符缓冲区
 * @return 32位无符号整数
 */
uint32_t ReadLittleEndianU32(const SmallVectorImpl<char> &CB) {
  assert(CB.size() == 4);
  return support::endian::read<uint32_t>(
      reinterpret_cast<const uint8_t *>(CB.data()), llvm::endianness::little);
}

/**
 * @brief 将ARM寄存器转换为Plan9寄存器格式
 *
 * @param armReg ARM寄存器名称
 * @return Plan9寄存器名称
 */
std::string ToPlan9Reg(const std::string &armReg) {
  if (armReg == "xzr" || armReg == "wzr") {
    return "ZR";
  }
  if (armReg == "lr") {
    return "R30";
  }
  std::string numPart;
  for (char c : armReg) {
    if (std::isdigit(c)) {
      numPart += c;
    }
  }
  if (numPart == "18") {
    return "R" + numPart + "_PLATFORM";
  }
  return "R" + numPart;
}

/**
 * @brief 将字符串转换为大写
 *
 * @param Str 输入字符串
 * @return 大写字符串
 */
std::string ToUpper(const std::string &Str) {
  std::string Res(Str.length(), 0);
  for (size_t I = 0; I < Str.length(); I++) {
    Res[I] = toUpper(Str[I]);
  }
  return Res;
};

/// 分支指令映射表
static std::unordered_map<std::string, std::string> BranchMap = {
    {"b", "B"},
    {"bl", "BL"},
    // isIndirectBranch，使用WORD指令
    // {"blr", "BLR"},
    // {"br", "BR"},

    {"b.eq", "BEQ"},
    {"b.ne", "BNE"},
    {"b.cs", "BCS"},
    {"b.hs", "BHS"}, // cs 同 hs
    {"b.cc", "BCC"},
    {"b.lo", "BLO"}, // cc 同 lo
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
    {"b.al", "B"}, // 总是真
    {"b.nv", "B"}, // b.nv 应该是不存在的

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
bool SLStreamer::MakeCmpareBranch(const std::vector<std::string> &Token,
                                  const std::string &InstStr) {
  auto &Op = Token[0];
  if (Op != "cbz" && Op != "cbnz" && Op != "tbz" && Op != "tbnz") {
    return false;
  }
  if (Op == "cbz") {
    this->Out << "    CMP ZR, " << ToPlan9Reg(Token[1]) << "\n";
    this->Out << "    BEQ ";
    tool::OutLabel(this->Out, Token[2]) << "  // " << InstStr << "\n";
  } else if (Op == "cbnz") {
    this->Out << "    CMP ZR, " << ToPlan9Reg(Token[1]) << "\n";
    this->Out << "    BNE ";
    tool::OutLabel(this->Out, Token[2]) << "  // " << InstStr << "\n";
  } else if (Op == "tbz") {
    this->Out << "    TST $(1<<" << Token[2].substr(1) << "), "
              << ToPlan9Reg(Token[1]) << "\n";
    this->Out << "    BEQ ";
    tool::OutLabel(this->Out, Token[3]) << "  // " << InstStr << "\n";
  } else {
    this->Out << "    TST $(1<<" << Token[2].substr(1) << "), "
              << ToPlan9Reg(Token[1]) << "\n";
    this->Out << "    BNE ";
    tool::OutLabel(this->Out, Token[3]) << "  // " << InstStr << "\n";
  }
  return true;
}

void SLStreamer::MakeBranch(const std::vector<std::string> &Token,
                            const std::string &InstStr) {
  auto &Op = Token[0];
  this->Out << "    " << BranchMap[Op] << " ";
  auto &Label = Token[1];
  tool::OutLabel(this->Out, Label);
  this->Out << "  // " << InstStr << "\n";
}

void SLStreamer::MakeBranchInst(const std::vector<std::string> &Token,
                                const std::string &InstStr) {
  auto &Op = Token[0];
  if (BranchMap.find(Op) != BranchMap.end()) {
    this->MakeBranch(Token, InstStr);
    this->ProgramCounter += 4;
    return;
  }
  if (this->MakeCmpareBranch(Token, InstStr)) {
    this->ProgramCounter += 8;
    return;
  }
  outs() << "Unsupported Branch Instruction\n";
}

void SLStreamer::emitInstruction(const MCInst &Inst,
                                 const MCSubtargetInfo &STI) {
  if (IsTopEmit == 0) {
    const auto &Desc = Bundle.getInstrInfo().get(Inst.getOpcode());
    std::string InstStr;
    raw_string_ostream OS(InstStr);
    Bundle.getInstPrinter().printInst(&Inst, 0, "", STI, OS);

    SmallVector<char> Buffer;
    SmallVector<MCFixup> Fixup;
    MCELFStreamer::getAssembler().getEmitter().encodeInstruction(
        Inst, Buffer, Fixup, Bundle.getSubtargetInfo());
    auto Token = tool::TokenizeInstruction(InstStr);
    // Fixup非空时，说明指令中存在需要在链接时处理的label参数
    // label参数在MCOperand中的判断是isExpr()，暂不清楚这种指令能否直接使用WORD表示
    if (Desc.isBranch() && !Desc.isIndirectBranch()) {
      this->MakeBranchInst(Token, InstStr);
    } else if (Token[0] == "adrp") {
      this->Out << "    ADR ";
      tool::OutLabel(this->Out, Token[2])
          << ", " << ToPlan9Reg(Token[1]) << "\n";
      this->ProgramCounter += 4;
    } else {
      this->Out << "    WORD $" << format_hex(ReadLittleEndianU32(Buffer), 10)
                << "  // " << InstStr << "\n";
      this->ProgramCounter += 4;
    }

    LLVM_DEBUG({
      dbgs() << "INSTRUCTION: " << InstStr << " Size:" << Buffer.size()
             << " Fixup:" << Fixup.empty() << "\n";
    });
  }
  IsTopEmit++;
  MCELFStreamer::emitInstruction(Inst, STI);
  IsTopEmit--;
}

void SLStreamer::MakeWordData(uint64_t Value, unsigned Size, unsigned Repeat) {
  // 将 Value 按 Size 字节小端写入缓冲区
  auto appendBytes = [&](uint64_t Val, unsigned S) {
    for (unsigned I = 0; I < S; ++I) {
      this->WordData += static_cast<char>(Val & 0xFF);
      Val >>= 8;
    }
  };

  // 写入 Repeat 次
  for (unsigned R = 0; R < Repeat; ++R) {
    appendBytes(Value, Size);
  }

  while (this->WordData.size() >= 4) {
    // 小端WORD
    uint32_t Word = (static_cast<uint8_t>(this->WordData[0]) << 0) |
                    (static_cast<uint8_t>(this->WordData[1]) << 8) |
                    (static_cast<uint8_t>(this->WordData[2]) << 16) |
                    (static_cast<uint8_t>(this->WordData[3]) << 24);

    Out << "    WORD $" << format_hex(Word, 10) << "\n";
    // 移除已输出的 4 字节
    this->WordData.erase(0, 4);
    this->ProgramCounter += 4;
  }
}

void SLStreamer::FlushPendingBytes() {
  if (this->WordData.empty()) {
    return;
  }
  // 补零到 4 字节对齐
  while (this->WordData.size() % 4 != 0) {
    this->WordData += '\0';
  }
  // 输出所有完整 WORD
  while (this->WordData.size() >= 4) {
    uint32_t Word = (static_cast<uint8_t>(this->WordData[0]) << 0) |
                    (static_cast<uint8_t>(this->WordData[1]) << 8) |
                    (static_cast<uint8_t>(this->WordData[2]) << 16) |
                    (static_cast<uint8_t>(this->WordData[3]) << 24);
    this->Out << "    WORD $" << format_hex(Word, 10) << "\n";
    WordData.erase(0, 4);
    this->ProgramCounter += 4;
  }
}

void SLStreamer::emitIntValue(uint64_t Value, unsigned Size) {
  if (IsTopEmit == 0) {
    this->MakeWordData(Value, Size);
    LLVM_DEBUG(dbgs() << "INT DATA: .int" << (Size * 8) << " = " << Value
                      << "\n");
  }
  IsTopEmit++;
  MCELFStreamer::emitIntValue(Value, Size);
  IsTopEmit--;
}

void SLStreamer::emitFill(const llvm::MCExpr &NumBytes, uint64_t FillValue,
                          llvm::SMLoc Loc) {
  if (IsTopEmit == 0) {
    int64_t NumBytesVal;
    bool Evaluated = NumBytes.evaluateAsAbsolute(NumBytesVal);
    this->MakeWordData(FillValue, 1, NumBytesVal);
    LLVM_DEBUG({
      dbgs() << "FILL (form 1): ";
      if (Evaluated) {
        dbgs() << ".space / .zero " << NumBytesVal
               << " bytes, fill value = " << FillValue;
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

void SLStreamer::emitBytes(StringRef Data) {
  if (IsTopEmit == 0) {
    this->WordData.append(Data.begin(), Data.end());
    while (this->WordData.size() >= 4) {
      uint32_t Word = (static_cast<uint8_t>(this->WordData[0]) << 0) |
                      (static_cast<uint8_t>(this->WordData[1]) << 8) |
                      (static_cast<uint8_t>(this->WordData[2]) << 16) |
                      (static_cast<uint8_t>(this->WordData[3]) << 24);
      this->Out << "    WORD $" << format_hex(Word, 10) << "\n";
      this->WordData.erase(0, 4);
      this->ProgramCounter += 4;
    }
    LLVM_DEBUG(dbgs() << "BYTE DATA: len=" << Data.size() << "\n");
  }
  IsTopEmit++;
  MCELFStreamer::emitBytes(Data);
  IsTopEmit--;
}

void SLStreamer::emitIdent(llvm::StringRef IdentString) {
  LLVM_DEBUG({ dbgs() << "IDENT: .ident \"" << IdentString << "\"\n"; });
  IsTopEmit++;
  MCELFStreamer::emitIdent(IdentString);
  IsTopEmit--;
}

void SLStreamer::emitValueToAlignment(llvm::Align Alignment, int64_t Value,
                                      unsigned ValueSize,
                                      unsigned MaxBytesToEmit) {
  LLVM_DEBUG(dbgs() << "Value Align: Alignment=" << Alignment.value()
                    << " Value=" << Value << " ValueSize=" << ValueSize
                    << " MaxBytesToEmit=" << MaxBytesToEmit << "\n");
  IsTopEmit++;
  MCELFStreamer::emitValueToAlignment(Alignment, Value, ValueSize,
                                      MaxBytesToEmit);
  IsTopEmit--;
}

void DumpDeclareHead(llvm::raw_fd_ostream &Out, const std::string &BaseName,
                     int64_t MaxDepth) {
  Out << "#include \"go_asm.h\"\n"
      << "#include \"funcdata.h\"\n"
      << "#include \"textflag.h\"\n";

  int64_t GoDepth = MaxDepth < 16 ? 0 : MaxDepth - 16;
  Out << "\nTEXT ·__" << BaseName << "_entry__(SB), NOSPLIT, $" << GoDepth
      << "\n"
      << "    NO_LOCAL_POINTERS\n"
      << "    WORD $0x100000a0 // adr x0, .+20\n"
      << "    MOVD R0, ret(FP)\n"
      << "    RET\n\n";
}

void DumpDeclareTail(llvm::raw_fd_ostream &Out, const std::string &BaseName,
                     tool::ParseResult &ParseRes, int64_t MaxDepth) {
  const auto &func = ParseRes.Funcs["__" + BaseName];
  if (!func.IsAllocated()) {
    outs() << BaseName << " IsAllocated(): false\n";
    return;
  }
  Out << "\nTEXT ·__" << BaseName << "(SB), NOSPLIT, $0-" << func.ArgSpace()
      << "\n"
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

  Out << "\n_" << BaseName << ":\n";
  size_t Offset = 0;
  for (auto &P : func.Params) {
    if (P.CReg.Name[0] == 'x') {
      Out << "    MOVD " << P.Name << "+" << Offset << "(FP), R"
          << P.CReg.Name.substr(1) << "\n";
    }
    if (P.CReg.Name[0] == 'd') {
      Out << "    FMOVD " << P.Name << "+" << Offset << "(FP), F"
          << P.CReg.Name.substr(1) << "\n";
    }
    Offset += P.Size;
  }
  Out << "    MOVD ·_subr__" << BaseName << "(SB), R11\n"
      << "    BL (R11)\n";

  if (!func.Results.empty()) {
    auto &P = func.Results[0];
    if (P.CReg.Name[0] == 'x') {
      Out << "    MOVD R0, " << P.Name << "+" << Offset << "(FP)\n";
    }
    if (P.CReg.Name[0] == 'd') {
      Out << "    FMOVD F0, " << P.Name << "+" << Offset << "(FP)\n";
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

void DumpSubrSL(const std::string &OutputPath, const std::string &Package,
                const std::string &BaseName, uint64_t StartPC,
                int64_t MaxDepth) {
  SmallString<256> SubrSL;
  sys::path::append(SubrSL, OutputPath,
                    (Twine(BaseName) + "_subr_arm64.go").str());
  std::error_code EC;
  raw_fd_ostream Out(SubrSL, EC, sys::fs::OF_None);
  if (EC) {
    outs() << EC.message() << "\n";
    return;
  }

  Out << "package " << Package << "\n\n";

  Out << "//go:nosplit\n"
      << "//go:noescape\n"
      << "//goland:noinspection ALL\n"
      << "func __" << BaseName << "_entry__() uintptr\n\n";

  Out << "var (\n"
      << "    _subr__" << BaseName << " uintptr = __" << BaseName
      << "_entry__() + " << StartPC << "\n"
      << ")\n\n";

  Out << "const (\n"
      << "    _stack__" << BaseName << " = " << MaxDepth << "\n"
      << ")\n\n";

  Out << "var (\n"
      << "    _ = _subr__" << BaseName << "\n"
      << ")\n\n";

  Out << "const (\n"
      << "    _ = _stack__" << BaseName << "\n"
      << ")\n\n";
}

} // namespace asm2arm
} // namespace tool