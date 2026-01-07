#ifndef DUMP_ELF_H
#define DUMP_ELF_H

#include "llvm/ADT/StringRef.h"
#include "llvm/MC/MCAsmBackend.h"

struct FuncRange {
    uint64_t StartAddr;
    uint64_t EndAddr;  // 左闭右开
    std::string Name;
};

void DumpElf(llvm::StringRef ElfPath);

#endif