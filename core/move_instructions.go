package core

/* MOV - Move (immediate)
 * ARM ARM A7.7.75 */
type MovImm InstrFields

func MovImm16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Imm := raw_instr & 0xff
	Rd := uint8((raw_instr >> 8) & 0x7)

	return MovImm{Rd: Rd, Rm: 0, Rn: 0, Imm: Imm, setflags: NOT_IT}
}

func (instr MovImm) Execute(regs *Registers) {
	value := instr.Imm

	MoveValue(regs, instr.Rd, value, instr.setflags, regs.Apsr.C)
}

/* MOV - Move (register)
 * ARM ARM A7.7.76
 * Encoding T1 */
type MovRegT1 InstrFields

func MovReg16T1(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Rd := uint8(raw_instr & 0x7)
	Rm := uint8((raw_instr >> 3) & 0xf)
	D := uint8((raw_instr >> 7) & 0x1)

	d := (D << 3) | Rd

	return MovRegT1{Rd: d, Rm: Rm, Rn: 0, Imm: 0, setflags: NEVER}
}

func (instr MovRegT1) Execute(regs *Registers) {
	if instr.Rd == 15 && regs.InITBlock() && !regs.LastInITBlock() {
		// UNPREDICTABLE
		// Raise exception (UsageFault?)
		return
	}

	MoveRegister(regs, instr.Rd, instr.Rm, instr.setflags, regs.Apsr.C)
}
