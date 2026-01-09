#ifndef UTILS_H
#define UTILS_H

#include "mc_bundle.h"

#include "llvm/MC/MCELFStreamer.h"
#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCCodeEmitter.h"

void FindSP(MCContextBundle &Bundle);

void PrintAArch64RegTable();

void PrintInstHelper(const llvm::MCInst &Inst, MCContextBundle &Bundle, uint64_t Addr);

bool StartWith(std::string_view Str, std::string_view Prefix);

std::string GetSourceName(llvm::StringRef Path);

#endif