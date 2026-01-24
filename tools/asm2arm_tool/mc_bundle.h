
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

#ifndef MC_BUNDLE_H
#define MC_BUNDLE_H

#include "llvm/MC/MCAsmInfo.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCDisassembler/MCDisassembler.h"
#include "llvm/MC/MCInstPrinter.h"
#include "llvm/MC/MCInstrInfo.h"
#include "llvm/MC/MCRegisterInfo.h"
#include "llvm/MC/MCSubtargetInfo.h"
#include "llvm/MC/TargetRegistry.h"
#include "llvm/TargetParser/Triple.h"

#include <memory>

namespace tool {
namespace mc {

/// MCContextBundle - A bundle of MC-related objects for a specific target.
/// This class manages the lifetime of various MC objects and provides
/// convenient access to them.
class MCContextBundle {
private:
  /// Triple - The target triple.
  const llvm::Triple &Triple;

  /// Target - The target object.
  const llvm::Target *Target;

  /// TargetOptions - The MC target options.
  llvm::MCTargetOptions TargetOptions;

  /// RegisterInfo - The MC register info.
  std::unique_ptr<llvm::MCRegisterInfo> RegisterInfo;

  /// AsmInfo - The MC assembly info.
  std::unique_ptr<llvm::MCAsmInfo> AsmInfo;

  /// SubtargetInfo - The MC subtarget info.
  std::unique_ptr<llvm::MCSubtargetInfo> SubtargetInfo;

  /// InstrInfo - The MC instruction info.
  std::unique_ptr<llvm::MCInstrInfo> InstrInfo;

  /// InstPrinter - The MC instruction printer.
  std::unique_ptr<llvm::MCInstPrinter> InstPrinter;

public:
  /// Constructor.
  ///
  /// @param TheTriple - The target triple.
  /// @param Features - The target features string.
  explicit MCContextBundle(const llvm::Triple &TheTriple,
                           const std::string &Features);

  /// Disable copy constructor and assignment operator.
  MCContextBundle(const MCContextBundle &) = delete;
  MCContextBundle &operator=(const MCContextBundle &) = delete;

  /// getRegisterInfo - Get the MC register info.
  const llvm::MCRegisterInfo &getRegisterInfo() { return *RegisterInfo; }

  /// getAsmInfo - Get the MC assembly info.
  const llvm::MCAsmInfo &getAsmInfo() { return *AsmInfo; }

  /// getSubtargetInfo - Get the MC subtarget info.
  const llvm::MCSubtargetInfo &getSubtargetInfo() { return *SubtargetInfo; }

  /// getInstrInfo - Get the MC instruction info.
  llvm::MCInstrInfo &getInstrInfo() { return *InstrInfo; }

  /// getInstPrinter - Get the MC instruction printer.
  llvm::MCInstPrinter &getInstPrinter() { return *InstPrinter; }

  /// getTarget - Get the target object.
  const llvm::Target &getTarget() { return *Target; }

  /// getTargetOptions - Get the MC target options.
  llvm::MCTargetOptions &getTargetOptions() { return TargetOptions; }

  /// getTriple - Get the target triple.
  const llvm::Triple &getTriple() { return Triple; }
};

} // end namespace mc
} // end namespace tool

#endif // MC_BUNDLE_H