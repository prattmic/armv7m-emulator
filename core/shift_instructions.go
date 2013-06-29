package core

/* LSL - Logical Shift Left (immediate)
 * ARM ARM A7.7.67 */
type LslImm InstrFields

func LslImm16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rd := uint8(raw_instr & 0x7)
	Rm := uint8((raw_instr >> 3) & 0x7)
	Imm := (raw_instr >> 6) & 0x1f

	return LslImm{Rd: Rd, Rm: Rm, Imm: Imm, S: true, Rn: 0}
}

func (instr LslImm) Execute(regs *Registers) {
	value := regs.R[instr.Rn]
	shift_n := uint8(instr.Imm)

	regs.R[instr.Rd] = LSL(regs, value, shift_n, instr.S)
}

/* LSL - Logical Shift Left (register)
 * ARM ARM A7.7.68 */
type LslReg InstrFields

func LslReg16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rdn := uint8(raw_instr & 0x7)
	Rm := uint8((raw_instr >> 3) & 0x7)

	return LslReg{Rd: Rdn, Rn: Rdn, Rm: Rm, Imm: 0, S: true}
}

func (instr LslReg) Execute(regs *Registers) {
	value := regs.R[instr.Rn]
	shift_n := uint8(regs.R[instr.Rm])

	regs.R[instr.Rd] = LSL(regs, value, shift_n, instr.S)
}

/* LSR - Logical Shift Right (immediate)
 * ARM ARM A7.7.69 */
type LsrImm InstrFields

func LsrImm16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rd := uint8(raw_instr & 0x7)
	Rm := uint8((raw_instr >> 3) & 0x7)
	Imm := (raw_instr >> 6) & 0x1f

	return LsrImm{Rd: Rd, Rm: Rm, Imm: Imm, S: true, Rn: 0}
}

func (instr LsrImm) Execute(regs *Registers) {
	regs.R[instr.Rd] = regs.R[instr.Rm] >> instr.Imm
}

/* LSR - Logical Shift Right (register)
 * ARM ARM A7.7.70 */
type LsrReg InstrFields

func LsrReg16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rdn := uint8(raw_instr & 0x7)
	Rm := uint8((raw_instr >> 3) & 0x7)

	return LsrReg{Rd: Rdn, Rn: Rdn, Rm: Rm, Imm: 0, S: true}
}

func (instr LsrReg) Execute(regs *Registers) {
	regs.R[instr.Rd] = regs.R[instr.Rn] >> uint8(regs.R[instr.Rm])
}
