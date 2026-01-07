#include "streamer.h"
#include "utils.h"

#include "llvm/MC/MCAssembler.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCObjectFileInfo.h"
#include "llvm/Support/raw_ostream.h"

using namespace llvm;

extern std::map<std::string, unsigned> AArch64RegTable;

PaddingNopObjectStreamer::PaddingNopObjectStreamer(llvm::MCContext &Context, std::unique_ptr<llvm::MCAsmBackend> TAB,
    std::unique_ptr<llvm::MCObjectWriter> OW, std::unique_ptr<llvm::MCCodeEmitter> Emitter,
    const llvm::MCInstrInfo *MCII, const llvm::MCRegisterInfo *MRI)
    : MCELFStreamer(Context, std::move(TAB), std::move(OW), std::move(Emitter)), MCII(MCII), MRI(MRI)
{
    FindSP(this->MRI);
    errs() << "AArch64RegTable[SP] = " << AArch64RegTable["SP"] << "\n";
}

void PaddingNopObjectStreamer::finish()
{
    this->PadTextSectionTo16Bytes();
    MCELFStreamer::finish();
}
void PaddingNopObjectStreamer::PadTextSectionTo16Bytes()
{
    MCSection *Text = this->getContext().getObjectFileInfo()->getTextSection();
    this->switchSection(Text);
    MCAssembler &Asm = this->getAssembler();
    uint64_t Size = Asm.getSectionAddressSize(*Text);
    uint64_t Pad = (16 - (Size % 16)) % 16;
    uint32_t AArch64NOP = 0xd503201f;
    while (Pad >= 4) {
        emitIntValue(AArch64NOP, 4);
        Pad -= 4;
    }
}