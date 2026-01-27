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

#include "dump_elf.h"
#include "utils.h"

#include "llvm/MC/MCAsmInfo.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCDisassembler/MCDisassembler.h"
#include "llvm/MC/MCInst.h"
#include "llvm/MC/MCInstrDesc.h"
#include "llvm/MC/MCSubtargetInfo.h"
#include "llvm/MC/TargetRegistry.h"
#include "llvm/Object/ELFObjectFile.h"
#include "llvm/Object/ObjectFile.h"
#include "llvm/Support/Debug.h"
#include "llvm/Support/FileSystem.h"
#include "llvm/Support/Format.h"
#include "llvm/Support/LineIterator.h"
#include "llvm/Support/MemoryBuffer.h"
#include "llvm/Support/NativeFormatting.h"
#include "llvm/Support/Path.h"
#include "llvm/Support/SourceMgr.h"
#include "llvm/Support/raw_ostream.h"

#include <cstdint>
#include <set>
#include <sys/types.h>
#include <unordered_map>

using namespace llvm;
using namespace llvm::object;
#define DEBUG_TYPE "dump_elf"

namespace tool {

namespace asm2arm {

/**
 * @brief 文本段指令列表
 *
 * 存储从ELF文件中反汇编得到的指令
 */
std::vector<MCInst> Text;

/**
 * @brief 指令地址列表
 *
 * 存储每条指令对应的地址
 */
std::vector<uint64_t> TextPC;

/**
 * @brief 指令大小列表
 *
 * 存储每条指令的大小（字节）
 */
std::vector<uint32_t> TextSize;

/**
 * @brief 地址到索引的映射
 *
 * 从指令地址映射到其在Text中的索引
 */
std::unordered_map<uint64_t, size_t> Addr2Idx;

/**
 * @brief 函数范围列表
 *
 * 存储从ELF文件中收集的函数范围信息
 */
std::vector<FuncRange> Funcs;

/**
 * @brief 返回指令地址集合
 *
 * 存储所有返回指令的地址
 */
std::set<uint64_t> Rets;

/**
 * @brief 收集函数范围信息
 *
 * 从ELF对象文件中收集函数的地址范围信息
 * @param Obj ELF对象文件
 */
static void CollectFuncRanges(ObjectFile &Obj) {
  auto *ELFObj = dyn_cast<ELFObjectFileBase>(&Obj);
  if (!ELFObj) {
    return;
  }
  for (const ELFSymbolRef &Sym : ELFObj->symbols()) {
    Expected<SymbolRef::Type> Type = Sym.getType();
    if (!Type || *Type != SymbolRef::ST_Function) {
      continue;
    }

    Expected<StringRef> Name = Sym.getName();
    Expected<uint64_t> Addr = Sym.getAddress();
    Expected<uint64_t> Size = Sym.getSize();
    if (!Name || !Addr || !Size) {
      continue;
    }

    FuncRange Func;
    Func.StartAddr = *Addr;
    Func.EndAddr = *Addr + *Size;
    Func.Name = Name->str();
    Funcs.push_back(Func);
    LLVM_DEBUG(dbgs() << Func.Name << " " << format_hex(Func.StartAddr, 6)
                      << " " << format_hex(Func.EndAddr, 6) << "\n";);
  }
  llvm::sort(Funcs, [](auto &A, auto &B) { return A.StartAddr < B.StartAddr; });
}

/**
 * @brief 反汇编文本段
 *
 * 反汇编ELF文件的文本段，提取指令信息
 * @param Bundle MC 上下文捆绑
 * @param Obj ELF对象文件
 * @param Sec 文本段
 * @return 反汇编是否成功
 */
static bool DisasmTextSection(tool::mc::MCContextBundle &Bundle,
                              const ObjectFile &Obj, const SectionRef &Sec) {
  Expected<StringRef> ContentExp = Sec.getContents();
  if (!ContentExp) {
    outs() << "getContents failed\n";
    return false;
  }

  StringRef Content = *ContentExp;
  uint64_t SectionAddr = Sec.getAddress();

  MCContext Ctx(Bundle.getTriple(), &Bundle.getAsmInfo(),
                &Bundle.getRegisterInfo(), &Bundle.getSubtargetInfo());
  std::unique_ptr<MCDisassembler> Disasm(
      Bundle.getTarget().createMCDisassembler(Bundle.getSubtargetInfo(), Ctx));
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

    auto DisasmStat =
        Disasm->getInstruction(Inst, InstSize, Bytes, CurAddr, errs());
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

/**
 * @brief 转储原始字节
 *
 * 将原始字节数据以十六进制格式转储到输出流
 * @param OS 输出流
 * @param Data 数据指针
 * @param Offset 数据偏移量
 * @param Size 数据大小
 */
static void DumpRawBytes(std::unique_ptr<raw_ostream> &OS, const uint8_t *Data,
                         size_t Offset, size_t Size) {
  for (size_t I = 0; I < Size;) {
    *OS << "    ";
    size_t LineEnd = std::min(I + 16, Size);
    for (size_t J = I; J < LineEnd; ++J) {
      *OS << format_hex(Data[Offset + J], 4) << ", ";
    }
    *OS << "   // data\n";
    I = LineEnd;
  }
}

void DumpElf(const std::string &OutputPath, StringRef ElfPath,
             tool::mc::MCContextBundle &Bundle, const std::string &Package,
             const std::string &BaseName, uint64_t &DumpTextSize,
             const std::string &Mode) {
  SmallString<256> DumpPath;
  sys::path::append(DumpPath, OutputPath,
                    (Twine(BaseName) + "_text_arm64.go").str());
  auto Buffer = MemoryBuffer::getFile(ElfPath);
  if (!Buffer) {
    outs() << "open ELF file failed\n";
    return;
  }
  Expected<std::unique_ptr<ObjectFile>> ObjectFileExp =
      ObjectFile::createObjectFile((*Buffer)->getMemBufferRef());
  if (!ObjectFileExp) {
    outs() << "createObjectFile failed\n";
    return;
  }
  ObjectFile &Object = **ObjectFileExp;
  CollectFuncRanges(Object); // 获取函数起止地址

  std::unique_ptr<raw_ostream> DumpStream;
  if (Mode == "JIT") {
    std::error_code EC;
    DumpStream =
        std::make_unique<raw_fd_ostream>(DumpPath, EC, sys::fs::OF_None);
    if (EC) {
      outs() << "Dump file error: " << EC.message() << "\n";
      return;
    }
  } else {
    DumpStream = std::make_unique<raw_null_ostream>();
  }

  *DumpStream << "package " << Package << "\n\n";
  *DumpStream << "var _text_" << BaseName << " = []byte{\n";

  for (auto &Sec : Object.sections()) {
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

    *DumpStream << "    // " << format_hex(BaseAddr, 18)
                << " Contents of section " << Name << "\n";

    if (Sec.isText()) {
      DisasmTextSection(Bundle, Object, Sec);

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
        DumpRawBytes(DumpStream, Bytes, ByteIndex, GapSize);
        ByteIndex += GapSize;
        CurrentAddr += GapSize;
      }

      for (size_t I = 0; I < NumInsts; ++I) {
        uint64_t InstAddr = TextPC[I];
        uint32_t InstLen = TextSize[I];

        assert(InstAddr == CurrentAddr && "Address misalignment!");
        assert(ByteIndex + InstLen <= TotalSize &&
               "Instruction overflows section");

        // 输出指令字节（单行）
        *DumpStream << "    ";
        for (uint32_t J = 0; J < InstLen; ++J) {
          *DumpStream << format_hex(Bytes[ByteIndex + J], 4) << ", ";
        }

        // 指令注释
        std::string InstStr;
        raw_string_ostream OSS(InstStr);
        Bundle.getInstPrinter().printInst(&Text[I], InstAddr, {},
                                          Bundle.getSubtargetInfo(), OSS);
        *DumpStream << "   // " << format_hex(InstAddr, 18) << " " << InstStr
                    << "\n";

        ByteIndex += InstLen;
        CurrentAddr += InstLen;

        // 计算到下一条指令（或段尾）的 gap
        uint64_t NextInstAddr =
            (I + 1 < NumInsts) ? TextPC[I + 1] : (BaseAddr + TotalSize);
        if (CurrentAddr < NextInstAddr) {
          size_t GapSize = NextInstAddr - CurrentAddr;
          DumpRawBytes(DumpStream, Bytes, ByteIndex, GapSize);
          ByteIndex += GapSize;
          CurrentAddr = NextInstAddr;
        }
      }
    } else if (Sec.isBSS()) {
      // .bss: 全零
      DumpTextSize += Size;
      for (uint64_t I = 0; I < Size; I += 16) {
        *DumpStream << "    ";
        uint64_t LineBytes = std::min<uint64_t>(16, Size - I);
        for (uint64_t J = 0; J < LineBytes; ++J) {
          *DumpStream << "0x00, ";
        }
        *DumpStream << "   \n";
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

      for (size_t I = 0; I < DataSize; I += 16) {
        *DumpStream << "    ";
        uint64_t LineBytes = std::min<uint64_t>(16, DataSize - I);
        for (size_t J = 0; J < LineBytes; ++J) {
          *DumpStream << format_hex(Data[I + J], 4) << ", ";
        }
        *DumpStream << "   \n";
      }
    }
  }
  *DumpStream << "}\n";
}

void DumpSubr(const BasicBlock &EntryBB, const std::string &Package,
              const std::string &OutputPath, const std::string &BaseName,
              const std::vector<std::pair<uint64_t, int64_t>> &SPDelta,
              const std::vector<int64_t> &Depth, uint64_t DumpTextSize) {
  SmallString<256> DumpPath;
  sys::path::append(DumpPath, OutputPath, (Twine(BaseName) + "_subr.go").str());
  std::error_code EC;
  raw_fd_ostream DumpStream(DumpPath, EC, sys::fs::OF_None);
  if (EC) {
    outs() << EC.message() << "\n";
    return;
  }

  DumpStream << "package " << Package << "\n\n"
             << "import (\n    `github.com/bytedance/sonic/loader`\n)\n\n"
             << "const (\n    _entry__" << BaseName << " = "
             << EntryBB.StartAddr << "\n)\n\n";

  int64_t MaxDepth = 0;
  for (auto X : Depth) {
    MaxDepth = std::max(MaxDepth, X);
  }
  DumpStream << "const (\n    _stack__" << BaseName << " = " << MaxDepth
             << "\n)\n\n"
             << "const (\n    _size__" << BaseName << " = " << DumpTextSize
             << "\n)\n\n"
             << "var (\n    _pcsp__" << BaseName << " = [][2]uint32{\n"
             << "        {0x1, 0},\n";

  std::map<uint64_t, int64_t> SPDump;
  for (size_t I = 1; I < SPDelta.size(); I++) {
    SPDump[SPDelta[I].first] = Depth[I];
  }
  for (auto X : Rets) {
    SPDump[X] = 0;
  }
  for (auto [Addr, SP] : SPDump) {
    DumpStream << "        {0x" << Twine::utohexstr(Addr) << ", " << SP
               << "},\n";
  }

  DumpStream << "        {0x" << Twine::utohexstr(DumpTextSize) << ", "
             << MaxDepth << "},\n"
             << "    }\n)\n\n"
             << "var _cfunc_" << BaseName << " = []loader.CFunc{\n    {\"_"
             << BaseName << "_entry\", 0, _entry__" << BaseName
             << ", 0, nil},\n"
             << "    {\"_" << BaseName << "\", _entry__" << BaseName
             << ", _size__" << BaseName << ", _stack__" << BaseName
             << ", _pcsp__" << BaseName << "},\n"
             << "}\n";
}

void DumpTmpl(const std::string &TmplFile, const std::string &Package,
              const std::string &OutputPath, const std::string &BaseName) {
  ErrorOr<std::unique_ptr<MemoryBuffer>> BufferOrErr =
      MemoryBuffer::getFile(TmplFile);
  if (std::error_code EC = BufferOrErr.getError()) {
    outs() << "Failed to open template file '" << TmplFile
           << "': " + EC.message() << "\n";
    return;
  }
  MemoryBuffer &Buffer = *BufferOrErr.get();

  SmallString<256> OutPath;
  sys::path::append(OutPath, OutputPath, (Twine(BaseName) + ".go").str());

  std::error_code EC;
  raw_fd_ostream OutFile(OutPath, EC, sys::fs::OF_Text);
  if (EC) {
    outs() << "Failed to create output file '" << OutPath.str()
           << "': " << EC.message() << "\n";
  }

  bool FoundPackageLine = false;
  StringRef Placeholder = "package {{PACKAGE}}";

  for (line_iterator LineIt(Buffer, false, '\0'); !LineIt.is_at_eof();
       ++LineIt) {
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

} // end namespace asm2arm

} // end namespace tool