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

class MCContextBundle {
public:
    explicit MCContextBundle(const llvm::Triple &TheTriple);

    MCContextBundle(const MCContextBundle &) = delete;
    MCContextBundle &operator=(const MCContextBundle &) = delete;

    const llvm::MCRegisterInfo &getMCRegisterInfo()
    {
        return *MRI;
    }

    const llvm::MCAsmInfo &getMCAsmInfo()
    {
        return *MAI;
    }

    const llvm::MCSubtargetInfo &getMCSubtargetInfo()
    {
        return *STI;
    }

    llvm::MCInstrInfo &getMCInstrInfo()
    {
        return *MCII;
    }

    llvm::MCInstPrinter &getMCInstPrinter()
    {
        return *MCIP;
    }

    const llvm::Target &getTarget()
    {
        return *TheTarget;
    }

    llvm::MCTargetOptions &getMCTargetOptions()
    {
        return MCOptions;
    }

    const llvm::Triple &getTriple()
    {
        return TheTriple;
    }

private:
    const llvm::Triple &TheTriple;
    const llvm::Target *TheTarget;
    llvm::MCTargetOptions MCOptions{};
    std::unique_ptr<llvm::MCRegisterInfo> MRI;
    std::unique_ptr<llvm::MCAsmInfo> MAI;
    std::unique_ptr<llvm::MCSubtargetInfo> STI;
    std::unique_ptr<llvm::MCInstrInfo> MCII;
    std::unique_ptr<llvm::MCInstPrinter> MCIP;
};
#endif