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

#include "utils.h"

#include "llvm/ADT/StringExtras.h"
#include "llvm/ADT/StringRef.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCInstrInfo.h"
#include "llvm/Support/Debug.h"
#include "llvm/Support/FileSystem.h"
#include "llvm/Support/Path.h"
#include "llvm/Support/raw_ostream.h"

#include <set>

using namespace llvm;
using namespace llvm::sys;

namespace tool {

namespace asm2arm {

/**
 * @brief AArch64 寄存器表
 *
 * 存储 AArch64 架构的寄存器名称到寄存器编号的映射
 */
std::map<std::string, unsigned> AArch64RegTable;

void FindSP(tool::mc::MCContextBundle &Bundle) {
  for (unsigned Reg = 0; Reg < Bundle.getRegisterInfo().getNumRegs(); Reg++) {
    AArch64RegTable[Bundle.getRegisterInfo().getName(Reg)] = Reg;
  }
  if (AArch64RegTable.find("SP") == AArch64RegTable.end()) {
    llvm::report_fatal_error("SP register not found!");
  }
}

void PrintAArch64RegTable() {
  for (auto &[Reg, Value] : AArch64RegTable) {
    dbgs() << "reg: " << Reg << " value: " << Value << "\n";
  }
}

void PrintInstHelper(const llvm::MCInst &Inst,
                     tool::mc::MCContextBundle &Bundle, uint64_t Addr) {
  dbgs() << "\n" << format_hex(Addr, 6) << "\n";
  StringRef Mnem = Bundle.getInstrInfo().getName(Inst.getOpcode());
  dbgs() << "Mnem=" << Mnem;
  unsigned NumOperands = Inst.getNumOperands();
  for (unsigned I = 0; I < NumOperands; I++) {
    dbgs() << " Operand" << std::to_string(I) << Inst.getOperand(I);
  }
  dbgs() << "\n";
  Inst.print(dbgs(), &Bundle.getRegisterInfo());
  dbgs() << "\n";

  const MCInstrDesc &Desc = Bundle.getInstrInfo().get(Inst.getOpcode());
  if (Desc.hasDefOfPhysReg(Inst, AArch64RegTable["SP"],
                           Bundle.getRegisterInfo())) {
    dbgs() << "修改了SP\n";
  }
  if (Desc.isPreISelOpcode()) {
    dbgs() << "前端伪指令\n";
  }
}

} // end namespace asm2arm

bool StartWith(std::string_view Str, std::string_view Prefix) {
  return Str.substr(0, Prefix.size()) == Prefix;
}

std::string GetSourceName(llvm::StringRef Path) {
  if (Path.empty()) {
    llvm::outs() << "error: empty file path\n";
    return "";
  }

  fs::file_status Status;
  if (fs::status(Path, Status)) {
    llvm::outs() << "error: cannot access file '" << Path << "'\n";
    return "";
  }

  if (!fs::is_regular_file(Status)) {
    llvm::outs() << "error: not a regular file: '" << Path << "'\n";
    return "";
  }

  std::string Ext = path::extension(Path).str();
  std::transform(Ext.begin(), Ext.end(), Ext.begin(),
                 [](unsigned char C) { return std::tolower(C); });

  static const std::set<std::string> ValidExts = {".s", ".S"};

  if (ValidExts.find(Ext) == ValidExts.end()) {
    llvm::outs() << "error: not a ASM file: '" << Path << "'\n";
    return "";
  }

  return path::stem(Path).str();
}

std::vector<std::string> TokenizeInstruction(const std::string &InstStr) {
  std::string Str = InstStr;
  // 去掉前导空白
  size_t Start = Str.find_first_not_of(" \t");
  if (Start == std::string::npos)
    return {};
  Str = Str.substr(Start);

  std::vector<std::string> Tokens;

  // 1. 提取操作码（直到空格）
  size_t Idx = 0;
  while (Idx < Str.size() && !isspace(Str[Idx])) {
    Idx++;
  }
  Tokens.push_back(Str.substr(0, Idx));
  if (Idx >= Str.size())
    return Tokens;

  // 2. 跳过空格
  while (Idx < Str.size() && isspace(Str[Idx]))
    Idx++;

  // 3. 解析操作数列表（支持 [x0, #8] 为一个整体）
  std::string CurrentToken;
  bool InBrackets = false;

  for (; Idx < Str.size(); ++Idx) {
    char Char = Str[Idx];

    if (Char == '[') {
      InBrackets = true;
      CurrentToken += Char;
    } else if (Char == ']') {
      InBrackets = false;
      CurrentToken += Char;
    } else if (Char == ',' && !InBrackets) {
      // 顶层逗号：结束当前操作数
      // 去掉尾部空格
      size_t End = CurrentToken.find_last_not_of(" \t");
      if (End != std::string::npos) {
        CurrentToken = CurrentToken.substr(0, End + 1);
      }
      Tokens.push_back(CurrentToken);
      CurrentToken.clear();
      // 跳过逗号后的空格
      while (Idx + 1 < Str.size() && isspace(Str[Idx + 1]))
        Idx++;
    } else {
      CurrentToken += Char;
    }
  }

  // 添加最后一个操作数
  if (!CurrentToken.empty()) {
    size_t End = CurrentToken.find_last_not_of(" \t");
    if (End != std::string::npos) {
      CurrentToken = CurrentToken.substr(0, End + 1);
    }
    Tokens.push_back(CurrentToken);
  }

  return Tokens;
}

llvm::raw_fd_ostream &OutLabel(llvm::raw_fd_ostream &Out,
                               llvm::StringRef Label) {
  for (auto &C : Label) {
    if (isAlpha(C) || isDigit(C) || C == '_') {
      Out << C;
    }
  }
  return Out;
}

} // end namespace tool