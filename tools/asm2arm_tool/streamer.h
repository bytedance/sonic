#ifndef STREAMER_H
#define STREAMER_H

#include "llvm/MC/MCELFStreamer.h"
#include "llvm/MC/MCInstrInfo.h"
#include "llvm/MC/MCAsmBackend.h"
#include "llvm/MC/MCObjectWriter.h"
#include "llvm/MC/MCCodeEmitter.h"
#include "llvm/MC/MCRegisterInfo.h"

#include <memory>

class PaddingNopObjectStreamer : public llvm::MCELFStreamer {
public:
    PaddingNopObjectStreamer(llvm::MCContext &Context, std::unique_ptr<llvm::MCAsmBackend> TAB,
        std::unique_ptr<llvm::MCObjectWriter> OW, std::unique_ptr<llvm::MCCodeEmitter> Emitter,
        const llvm::MCInstrInfo *MCII, const llvm::MCRegisterInfo *MRI);

    void finish();

protected:
    void PadTextSectionTo16Bytes();

    const llvm::MCRegisterInfo *MRI;
    const llvm::MCInstrInfo *MCII;
};

#endif