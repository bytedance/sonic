#ifndef CAL_DEPTH_H
#define CAL_DEPTH_H

#include "mc_bundle.h"

#include <cstdint>

struct BasicBlock {
    uint64_t StartAddr;
    uint64_t EndAddr;  // 左闭右闭
    size_t FirstIdx;   // Text下标
    size_t LastIdx;
};

struct BBSP {
    int64_t NetDelta = 0;  // SP总变化字节
    int64_t Peak = 0;      // 峰值 <= 0
};

struct BBPrefix {
    std::vector<int64_t> PreInstSP;  // BB块内SP变化前缀和
};

int SplitBasicBlocks(MCContextBundle &Bundle, std::vector<BasicBlock> &BBs, const std::string &EntryBB);

void BuildCFG(MCContextBundle &Bundle, std::vector<BasicBlock> &BBs, std::vector<std::vector<size_t>> &CFG);

bool HasCycle(const std::vector<BasicBlock> &BBs, const std::vector<std::vector<size_t>> &CFG);

void CalcSPDelta(MCContextBundle &Bundle, const std::vector<BasicBlock> &BBs, std::vector<BBSP> &BBSPVec,
    std::vector<std::pair<uint64_t, int64_t>> &SPDelta);

std::vector<int64_t> ComputeMaxSPDepth(const std::vector<std::vector<size_t>> &CFG, const std::vector<BBSP> &BBSPVec);

std::vector<int64_t> ComputeMaxSPDepthAtInst(size_t entry_bb, const std::vector<BasicBlock> &BBs,
    const std::vector<std::vector<size_t>> &CFG, const std::vector<BBSP> &BBSPVec,
    const std::vector<std::pair<uint64_t, int64_t>> &SPDelta);

#endif