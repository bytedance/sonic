#ifndef DUMP_ELF_H
#define DUMP_ELF_H

#include "mc_bundle.h"

#include "llvm/ADT/StringRef.h"
#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCInstPrinter.h"

struct FuncRange {
    uint64_t StartAddr;
    uint64_t EndAddr;  // 左闭右开
    std::string Name;
};

void DumpElf(llvm::StringRef ElfPath, llvm::StringRef DumpFile, MCContextBundle &Bundle, const std::string &Package,
    const std::string &BaseName);

#endif