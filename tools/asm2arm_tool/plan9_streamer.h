#ifndef PLAN9_STREAMER_H
#define PLAN9_STREAMER_H

#include "mc_bundle.h"

#include "llvm/MC/MCELFStreamer.h"
#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCObjectWriter.h"
#include "llvm/MC/MCCodeEmitter.h"
#include "llvm/MC/MCContext.h"
#include <cstddef>

/// adrp

/// 1、switchSection切换段时，plan9需要知道吗？我看好像已有的都没管这个
/// 2、emitSymbolAttribute和emitSymbolDesc，需要通过这个来加plan9的文件头尾吗
/// 3、代码、数据对齐，细节如何处理
///     3.1 arm64代码段肯定是4字节对齐
///     3.2 数据段在每个标签对应的数据大小后补齐至4字节对齐
///     3.3 .align N 或 .p2align N 可以直接丢掉

class Plan9Streamer : public llvm::MCELFStreamer {
public:
    Plan9Streamer(llvm::MCContext &Context, std::unique_ptr<llvm::MCAsmBackend> TAB,
        std::unique_ptr<llvm::MCObjectWriter> OW, std::unique_ptr<llvm::MCCodeEmitter> Emitter,
        llvm::raw_fd_ostream &Out, MCContextBundle &Bundle);

    // 获取label
    void emitLabel(llvm::MCSymbol *Sym, llvm::SMLoc Loc = {}) override;

    // 获取Instruction
    void emitInstruction(const llvm::MCInst &Inst, const llvm::MCSubtargetInfo &STI) override;

    // .ident 不用处理
    void emitIdent(llvm::StringRef IdentString) override;

    // .byte | .short | .hword | .word | .quad | .dword | .xword
    void emitIntValue(uint64_t Value, unsigned Size) override;

    // .fill repeat, size, value
    // 这个会调用 emitIntValue
    // void emitFill(const llvm::MCExpr &NumValues, int64_t Size, int64_t Expr, llvm::SMLoc Loc = llvm::SMLoc())
    // override;

    // .zero N |
    void emitFill(const llvm::MCExpr &NumBytes, uint64_t FillValue, llvm::SMLoc Loc = llvm::SMLoc()) override;

    // .asciz "string" | .ascii "string"
    // .asciz 会被拆开调用两次emitBytes，第二次是单独的0
    void emitBytes(llvm::StringRef Data) override;

    // .align N 或 .p2align N 不用处理
    void emitValueToAlignment(
        llvm::Align Alignment, int64_t Value = 0, unsigned ValueSize = 1, unsigned MaxBytesToEmit = 0) override;

    // 这个会调用emitValueToAlignment
    // void emitCodeAlignment(
    //         llvm::Align Alignment, const llvm::MCSubtargetInfo *STI, unsigned MaxBytesToEmit = 0) override;

    /*
        // .globl / .global
        bool emitSymbolAttribute(llvm::MCSymbol *Sym, llvm::MCSymbolAttr Attribute) override;

        // .type symbol, @function 或 @object
        void emitSymbolDesc(llvm::MCSymbol *Sym, unsigned DescValue) override;

        // .size symbol, expression
        void emitELFSize(llvm::MCSymbol *Sym, const llvm::MCExpr *Value) override;
    */
private:
    MCContextBundle &Bundle;
    size_t IsTopEmit = 0;
    llvm::raw_fd_ostream &Out;
    std::string WordData;

public:
    void makeWordData(uint64_t Value, unsigned Size, unsigned Repeat = 1);
    void flushPendingBytes();
    void makeBranchInst(const std::vector<std::string> &Token, const std::string &InstStr);
    void makeBranch(const std::vector<std::string> &Token, const std::string &InstStr);
    bool makeCmpareBranch(const std::vector<std::string> &Token, const std::string &InstStr);
    void finish();
};

#endif