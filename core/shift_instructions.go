package core

import "fmt"

/* LSL - Logical Shift Left (immediate)
 * ARM ARM A7.7.67 */
type LslImm InstrFields

func LslImm16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rd := RegIndex(raw_instr & 0x7)
	Rm := RegIndex((raw_instr >> 3) & 0x7)
	Imm := (raw_instr >> 6) & 0x1f

	if Imm == 0 {
		/* Equivalent to MOV (reg) T2 encoding */
		return MovReg16T2(instr)
	}

	return LslImm{Rd: Rd, Rm: Rm, Rn: 0, Imm: Imm, setflags: NOT_IT}
}

func (instr LslImm) Execute(regs *Registers) {
	value := regs.R(instr.Rm)
	shift_n := uint8(instr.Imm)

	result := LSL(regs, value, shift_n, instr.setflags)
	regs.SetR(instr.Rd, result)
}

func (instr LslImm) String() string {
	return fmt.Sprintf("lsl%s %s, %s, #%#x", instr.setflags, instr.Rd, instr.Rm, instr.Imm)
}

/* LSL - Logical Shift Left (register)
 * ARM ARM A7.7.68 */
type LslReg InstrFields

func LslReg16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rdn := RegIndex(raw_instr & 0x7)
	Rm := RegIndex((raw_instr >> 3) & 0x7)

	return LslReg{Rd: Rdn, Rn: Rdn, Rm: Rm, Imm: 0, setflags: NOT_IT}
}

func (instr LslReg) Execute(regs *Registers) {
	value := regs.R(instr.Rn)
	shift_n := uint8(regs.R(instr.Rm))

	result := LSL(regs, value, shift_n, instr.setflags)
	regs.SetR(instr.Rd, result)
}

func (instr LslReg) String() string {
	return fmt.Sprintf("lsl%s %s, %s", instr.setflags, instr.Rd, instr.Rm)
}

/* LSR - Logical Shift Right (immediate)
 * ARM ARM A7.7.69 */
type LsrImm InstrFields

func LsrImm16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rd := RegIndex(raw_instr & 0x7)
	Rm := RegIndex((raw_instr >> 3) & 0x7)
	Imm := (raw_instr >> 6) & 0x1f

	return LsrImm{Rd: Rd, Rm: Rm, Rn: 0, Imm: Imm, setflags: NOT_IT}
}

func (instr LsrImm) Execute(regs *Registers) {
	value := regs.R(instr.Rm)
	shift_n := uint8(instr.Imm)

	result := LSR(regs, value, shift_n, instr.setflags)
	regs.SetR(instr.Rd, result)
}

func (instr LsrImm) String() string {
	return fmt.Sprintf("lsr%s %s, %s, #%#x", instr.setflags, instr.Rd, instr.Rm, instr.Imm)
}

/* LSR - Logical Shift Right (register)
 * ARM ARM A7.7.70 */
type LsrReg InstrFields

func LsrReg16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rdn := RegIndex(raw_instr & 0x7)
	Rm := RegIndex((raw_instr >> 3) & 0x7)

	return LsrReg{Rd: Rdn, Rn: Rdn, Rm: Rm, Imm: 0, setflags: NOT_IT}
}

func (instr LsrReg) Execute(regs *Registers) {
	value := regs.R(instr.Rn)
	shift_n := uint8(regs.R(instr.Rm))

	result := LSR(regs, value, shift_n, instr.setflags)
	regs.SetR(instr.Rd, result)
}

func (instr LsrReg) String() string {
	return fmt.Sprintf("lsr%s %s, %s", instr.setflags, instr.Rd, instr.Rm)
}
