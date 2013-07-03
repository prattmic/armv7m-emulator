package core

import "fmt"

/* ADD (register)
 * ARM ARM A7.7.4
 * Encoding T1 */
type AddRegT1 InstrFields

func AddReg16T1(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rd := RegIndex(raw_instr & 0x7)
	Rn := RegIndex((raw_instr >> 3) & 0x7)
	Rm := RegIndex((raw_instr >> 6) & 0x7)

	return AddRegT1{Rd: Rd, Rm: Rm, Rn: Rn, Imm: 0, setflags: NOT_IT}
}

func (instr AddRegT1) Execute(regs *Registers) {
	AddRegister(regs, InstrFields(instr), Shift{function: LSL_C, amount: 0})
}

func (instr AddRegT1) String() string {
	return fmt.Sprintf("adds %s, %s, %s", instr.Rd, instr.Rn, instr.Rm)
}
