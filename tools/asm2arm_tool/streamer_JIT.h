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

#ifndef STREAMER_JIT_H
#define STREAMER_JIT_H

#include "mc_bundle.h"

#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCCodeEmitter.h"
#include "llvm/MC/MCELFStreamer.h"
#include "llvm/MC/MCObjectWriter.h"
#include "llvm/MC/MCRegisterInfo.h"

#include <memory>

namespace tool {
namespace asm2arm {

/// JITStreamer - An MCELFStreamer that pads the text section to
/// 16-byte alignment with NOP instructions for JIT compilation.
class JITStreamer : public llvm::MCELFStreamer {
private:
  /// MRI - The MCRegisterInfo instance.
  const llvm::MCRegisterInfo *MRI;

protected:
  /// PadTextSectionTo16Bytes - Pad the text section to 16-byte alignment using
  /// NOP instructions.
  void PadTextSectionTo16Bytes();

public:
  /// Constructor.
  ///
  /// @param Context - The MCContext instance.
  /// @param AsmBackend - The MCAsmBackend instance.
  /// @param ObjWriter - The MCObjectWriter instance.
  /// @param CodeEmitter - The MCCodeEmitter instance.
  /// @param Bundle - The MCContextBundle instance.
  JITStreamer(llvm::MCContext &Context,
              std::unique_ptr<llvm::MCAsmBackend> AsmBackend,
              std::unique_ptr<llvm::MCObjectWriter> ObjWriter,
              std::unique_ptr<llvm::MCCodeEmitter> CodeEmitter,
              tool::mc::MCContextBundle &Bundle);

  /// finish - Finish streamer operations.
  void finish();
};

} // end namespace asm2arm
} // end namespace tool

#endif // STREAMER_JIT_H
