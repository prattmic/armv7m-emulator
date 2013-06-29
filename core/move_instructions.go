package core

/* MOV - Move (immediate)
 * ARM ARM A7.7.75 */
type MovImm InstrFields

func MovImm16(instr FetchedInstr) DecodedInstr {
	raw_instr := instr.Uint32()

	Imm := raw_instr & 0xff
	Rd := uint8((raw_instr >> 8) & 0x7)

	return MovImm{Rd: Rd, Rm: 0, Rn: 0, Imm: Imm, S: true}
}

func (instr MovImm) Execute(regs *Registers) {
	result := instr.Imm

	regs.R[instr.Rd] = result

	if instr.S {
		regs.Apsr.N = (result & 0x80000000) != 0
		regs.Apsr.Z = result == 0
	}
}
