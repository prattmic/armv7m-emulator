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

/* ADD (register)
 * ARM ARM A7.7.4
 * Encoding T2 */
type AddRegT2 InstrFields

func AddReg16T2(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	rdn := uint8(raw_instr & 0x7)
	DN := uint8((raw_instr >> 7) & 0x1)
	Rdn := RegIndex((DN << 3) | rdn)
	Rm := RegIndex((raw_instr >> 3) & 0xf)

	if Rm == SP {
		return AddRegSP16T1(instr)
	}

	if Rdn == SP {
		return AddRegSP16T2(instr)
	}

	return AddRegT2{Rd: Rdn, Rm: Rm, Rn: Rdn, Imm: 0, setflags: NEVER}
}

func (instr AddRegT2) Execute(regs *Registers) {
	if instr.Rd == PC && regs.InITBlock() && !regs.LastInITBlock() {
		// UNPREDICTABLE
		// Raise exception (UsageFault?)
		return
	} else if instr.Rd == PC && instr.Rm == PC {
		// UNPREDICTABLE
		// Raise exception (UsageFault?)
		return
	}

	AddRegister(regs, InstrFields(instr), Shift{function: LSL_C, amount: 0})
}

func (instr AddRegT2) String() string {
	return fmt.Sprintf("add %s, %s", instr.Rd, instr.Rm)
}

/* ADD (SP plus register)
 * ARM ARM A7.7.6
 * Encoding T1 */
type AddRegSPT1 InstrFields

func AddRegSP16T1(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	rdm := uint8(raw_instr & 0x7)
	DM := uint8((raw_instr >> 7) & 0x1)

	Rdm := RegIndex((DM << 3) | rdm)

	return AddRegSPT1{Rd: Rdm, Rm: Rdm, Rn: SP, Imm: 0, setflags: NEVER}
}

func (instr AddRegSPT1) Execute(regs *Registers) {
	AddRegister(regs, InstrFields(instr), Shift{function: LSL_C, amount: 0})
}

func (instr AddRegSPT1) String() string {
	return fmt.Sprintf("add %s, sp, %s", instr.Rd, instr.Rd)
}

/* ADD (SP plus register)
 * ARM ARM A7.7.6
 * Encoding T2 */
type AddRegSPT2 InstrFields

func AddRegSP16T2(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rm := RegIndex((raw_instr >> 3) & 0xf)

	if Rm == SP {
		return AddRegSP16T1(instr)
	}

	return AddRegSPT2{Rd: SP, Rm: Rm, Rn: SP, Imm: 0, setflags: NEVER}
}

func (instr AddRegSPT2) Execute(regs *Registers) {
	AddRegister(regs, InstrFields(instr), Shift{function: LSL_C, amount: 0})
}

func (instr AddRegSPT2) String() string {
	return fmt.Sprintf("add sp, %s", instr.Rm)
}

/* SUB (register)
 * ARM ARM A7.7.172
 * Encoding T1 */
type SubRegT1 InstrFields

func SubReg16T1(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rd := RegIndex(raw_instr & 0x7)
	Rn := RegIndex((raw_instr >> 3) & 0x7)
	Rm := RegIndex((raw_instr >> 6) & 0x7)

	return SubRegT1{Rd: Rd, Rm: Rm, Rn: Rn, Imm: 0, setflags: NOT_IT}
}

func (instr SubRegT1) Execute(regs *Registers) {
	SubRegister(regs, InstrFields(instr), Shift{function: LSL_C, amount: 0})
}

func (instr SubRegT1) String() string {
	return fmt.Sprintf("subs %s, %s, %s", instr.Rd, instr.Rm, instr.Rn)
}
