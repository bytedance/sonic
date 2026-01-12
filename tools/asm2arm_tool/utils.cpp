#include "utils.h"

#include "llvm/ADT/StringRef.h"
#include "llvm/MC/MCContext.h"
#include "llvm/MC/MCInstrInfo.h"
#include "llvm/Support/Debug.h"
#include "llvm/Support/raw_ostream.h"
#include "llvm/Support/FileSystem.h"
#include "llvm/Support/Path.h"

#include <set>

using namespace llvm;
using namespace llvm::sys;

std::map<std::string, unsigned> AArch64RegTable;

void FindSP(MCContextBundle &Bundle)
{
    for (unsigned r = 0; r < Bundle.getMCRegisterInfo().getNumRegs(); r++) {
        AArch64RegTable[Bundle.getMCRegisterInfo().getName(r)] = r;
    }
    if (AArch64RegTable.find("SP") == AArch64RegTable.end()) {
        llvm::report_fatal_error("SP register not found!");
    }
}

void PrintAArch64RegTable()
{
    for (auto &[reg, v] : AArch64RegTable) {
        dbgs() << "reg: " << reg << " value: " << v << "\n";
    }
}

void PrintInstHelper(const llvm::MCInst &Inst, MCContextBundle &Bundle, uint64_t Addr)
{
    dbgs() << "\n" << format_hex(Addr, 6) << "\n";
    StringRef Mnem = Bundle.getMCInstrInfo().getName(Inst.getOpcode());
    dbgs() << "Mnem=" << Mnem;
    unsigned NumOperands = Inst.getNumOperands();
    for (unsigned i = 0; i < NumOperands; i++) {
        dbgs() << " Operand" << std::to_string(i) << Inst.getOperand(i);
    }
    dbgs() << "\n";
    Inst.print(dbgs(), &Bundle.getMCRegisterInfo());
    dbgs() << "\n";

    const MCInstrDesc &Desc = Bundle.getMCInstrInfo().get(Inst.getOpcode());
    if (Desc.hasDefOfPhysReg(Inst, AArch64RegTable["SP"], Bundle.getMCRegisterInfo())) {
        dbgs() << "修改了SP\n";
    }
    if (Desc.isPreISelOpcode()) {
        dbgs() << "前端伪指令\n";
    }
}

bool StartWith(std::string_view Str, std::string_view Prefix)
{
    return Str.substr(0, Prefix.size()) == Prefix;
}

std::string GetSourceName(llvm::StringRef Path)
{
    if (Path.empty()) {
        llvm::outs() << "error: empty file path\n";
        return "";
    }

    fs::file_status Status;
    if (fs::status(Path, Status)) {
        llvm::outs() << "error: cannot access file '" << Path << "'\n";
        return "";
    }

    if (!fs::is_regular_file(Status)) {
        llvm::outs() << "error: not a regular file: '" << Path << "'\n";
        return "";
    }

    std::string ext = path::extension(Path).str();
    std::transform(ext.begin(), ext.end(), ext.begin(), [](unsigned char c) { return std::tolower(c); });

    static const std::set<std::string> ValidExts = {".s", ".S"};

    if (ValidExts.find(ext) == ValidExts.end()) {
        llvm::outs() << "error: not a ASM file: '" << Path << "'\n";
        return "";
    }

    return path::stem(Path).str();
}