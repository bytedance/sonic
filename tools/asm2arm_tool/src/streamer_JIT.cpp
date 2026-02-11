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

#include "streamer_JIT.h"
#include "utils.h"

#include "llvm/MC/MCAssembler.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCAsmLayout.h"
#include "llvm/MC/MCObjectFileInfo.h"
#include "llvm/Support/Debug.h"
#include "llvm/Support/raw_ostream.h"

using namespace llvm;
#define DEBUG_TYPE "streamer"

namespace tool {
namespace asm2arm {

extern std::map<std::string, unsigned> AArch64RegTable;

JITStreamer::JITStreamer(MCContext &Context,
                         std::unique_ptr<MCAsmBackend> AsmBackend,
                         std::unique_ptr<MCObjectWriter> ObjWriter,
                         std::unique_ptr<MCCodeEmitter> CodeEmitter,
                         tool::mc::MCContextBundle &Bundle)
    : MCELFStreamer(Context, std::move(AsmBackend), std::move(ObjWriter),
                    std::move(CodeEmitter)),
      MRI(&Bundle.getRegisterInfo()) {
  LLVM_DEBUG(dbgs() << "AArch64RegTable[SP] = " << AArch64RegTable["SP"]
                    << "\n");
}

void JITStreamer::finish() {
  PadTextSectionTo16Bytes();
  MCELFStreamer::finish();
}

void JITStreamer::PadTextSectionTo16Bytes() {
  MCSection *TextSection = getContext().getObjectFileInfo()->getTextSection();
  switchSection(TextSection);

  MCAssembler &Assembler = getAssembler();
  MCAsmLayout ASMLayout = MCAsmLayout(Assembler);
  uint64_t SectionSize = ASMLayout.getSectionAddressSize(TextSection);
  uint64_t PaddingSize = (16 - (SectionSize % 16)) % 16;

  // AArch64 NOP instruction encoding
  const uint32_t AArch64NOP = 0xd503201f;

  while (PaddingSize >= 4) {
    emitIntValue(AArch64NOP, 4);
    PaddingSize -= 4;
  }
}

} // end namespace asm2arm
} // end namespace tool