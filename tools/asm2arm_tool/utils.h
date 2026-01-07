#ifndef UTILS_H
#define UTILS_H

#include "llvm/MC/MCELFStreamer.h"
#include "llvm/MC/MCInstrInfo.h"
#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCCodeEmitter.h"
#include "llvm/MC/MCRegisterInfo.h"

void FindSP(const llvm::MCRegisterInfo *MRI);

void PrintAArch64RegTable();

void PrintInstHelper(
    const llvm::MCInst &Inst, const llvm::MCRegisterInfo *MRI, const llvm::MCInstrInfo *MCII, uint64_t Addr);

bool StartWith(std::string_view Str, std::string_view Prefix);

#endif