#include "utils.h"

#include "llvm/ADT/StringRef.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCInstrInfo.h"
#include "llvm/Support/raw_ostream.h"

using namespace llvm;

std::map<std::string, unsigned> AArch64RegTable;

void FindSP(const llvm::MCRegisterInfo *MRI)
{
    for (unsigned r = 0; r < MRI->getNumRegs(); r++) {
        AArch64RegTable[MRI->getName(r)] = r;
    }
    if (AArch64RegTable.find("SP") == AArch64RegTable.end()) {
        llvm::report_fatal_error("SP register not found!");
    }
}

void PrintAArch64RegTable()
{
    for (auto &[reg, v] : AArch64RegTable) {
        errs() << "reg: " << reg << " value: " << v << "\n";
    }
}

void PrintInstHelper(
    const llvm::MCInst &Inst, const llvm::MCRegisterInfo *MRI, const llvm::MCInstrInfo *MCII, uint64_t Addr)
{
    errs() << "\n" << format_hex(Addr, 6) << "\n";
    StringRef Mnem = MCII->getName(Inst.getOpcode());
    errs() << "Mnem=" << Mnem;
    unsigned NumOperands = Inst.getNumOperands();
    for (unsigned i = 0; i < NumOperands; i++) {
        errs() << " Operand" << std::to_string(i) << Inst.getOperand(i);
    }
    errs() << "\n";
    Inst.print(errs(), MRI);
    errs() << "\n";

    const MCInstrDesc &Desc = MCII->get(Inst.getOpcode());
    if (Desc.hasDefOfPhysReg(Inst, AArch64RegTable["SP"], *MRI)) {
        errs() << "修改了SP\n";
    }
    if (Desc.isPreISelOpcode()) {
        errs() << "前端伪指令\n";
    }
}

bool StartWith(std::string_view Str, std::string_view Prefix)
{
    return Str.substr(0, Prefix.size()) == Prefix;
}