#include "mc_bundle.h"

#include "llvm/MC/TargetRegistry.h"

#include <cassert>

MCContextBundle::MCContextBundle(const llvm::Triple &TheTriple, const std::string &Features) : TheTriple(TheTriple)
{
    using namespace llvm;

    std::string Error;
    TheTarget = TargetRegistry::lookupTarget(TheTriple.getTriple(), Error);
    if (!TheTarget) {
        errs() << "MCContextBundle: " << Error << "\n";
        assert(false && "Target not found");
    }

    MRI.reset(TheTarget->createMCRegInfo(TheTriple.str()));
    assert(MRI && "Unable to create MCRegisterInfo!");

    MAI.reset(TheTarget->createMCAsmInfo(*MRI, TheTriple.str(), MCOptions));
    assert(MAI && "Unable to create MCAsmInfo!");

    STI.reset(TheTarget->createMCSubtargetInfo(TheTriple.str(), "generic", Features));
    assert(STI && "Unable to create MCSubtargetInfo!");

    MCII.reset(TheTarget->createMCInstrInfo());
    assert(MCII && "Unable to create MCInstrInfo!");

    MCIP.reset(TheTarget->createMCInstPrinter(TheTriple, MAI->getAssemblerDialect(), *MAI, *MCII, *MRI));
    assert(MCIP && "Unable to create MCInstPrinter!");
}