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

#include "mc_bundle.h"

#include "llvm/MC/TargetRegistry.h"
#include "llvm/Support/raw_ostream.h"

#include <cassert>

using namespace llvm;

namespace tool {
namespace mc {

MCContextBundle::MCContextBundle(const llvm::Triple &TheTriple,
                                 const std::string &Features)
    : Triple(TheTriple) {
  std::string Error;
  Target = TargetRegistry::lookupTarget(TheTriple.getTriple(), Error);
  if (!Target) {
    outs() << "MCContextBundle: " << Error << "\n";
    assert(false && "Target not found");
  }

  RegisterInfo.reset(Target->createMCRegInfo(TheTriple.str()));
  assert(RegisterInfo && "Unable to create MCRegisterInfo!");

  AsmInfo.reset(
      Target->createMCAsmInfo(*RegisterInfo, TheTriple.str(), TargetOptions));
  assert(AsmInfo && "Unable to create MCAsmInfo!");

  SubtargetInfo.reset(
      Target->createMCSubtargetInfo(TheTriple.str(), "generic", Features));
  assert(SubtargetInfo && "Unable to create MCSubtargetInfo!");

  InstrInfo.reset(Target->createMCInstrInfo());
  assert(InstrInfo && "Unable to create MCInstrInfo!");

  InstPrinter.reset(
      Target->createMCInstPrinter(TheTriple, AsmInfo->getAssemblerDialect(),
                                  *AsmInfo, *InstrInfo, *RegisterInfo));
  assert(InstPrinter && "Unable to create MCInstPrinter!");
}

} // end namespace mc
} // end namespace tool