#ifndef DUMP_ELF_H
#define DUMP_ELF_H

#include "mc_bundle.h"
#include "cal_depth.h"

#include "llvm/ADT/StringRef.h"
#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCInstPrinter.h"

struct FuncRange {
    uint64_t StartAddr;
    uint64_t EndAddr;  // 左闭右开
    std::string Name;
};

void DumpElf(const std::string &OutputPath, llvm::StringRef ElfPath, MCContextBundle &Bundle,
    const std::string &Package, const std::string &BaseName, uint64_t &DumpTextSize);

void DumpSubr(const BasicBlock &EntryBB, const std::string &Package, const std::string &OutputPath,
    const std::string &BaseName, const std::vector<std::pair<uint64_t, int64_t>> &SPDelta,
    const std::vector<int64_t> &Depth, uint64_t DumpTextSize);

void DumpTmpl(
    const std::string &TmplDir, const std::string &Package, const std::string &OutputPath, const std::string &BaseName);

#endif