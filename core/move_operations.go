package core

/* Move value into destination register, updating condition codes */
func MoveValue(regs *Registers, dest uint8, value uint32, setflags SetFlags, carry bool) {
	regs.R[dest] = value

	if setflags == ALWAYS || (setflags == NOT_IT && !regs.InITBlock()) {
		regs.Apsr.N = (value & 0x80000000) != 0
		regs.Apsr.Z = value == 0
		regs.Apsr.C = carry
	}
}

func MoveRegister(regs *Registers, dest uint8, source uint8, setflags SetFlags, carry bool) {
	value := regs.R[source]

	if dest == PC {
		regs.ALUWritePC(value)
	} else {
		MoveValue(regs, dest, value, setflags, carry)
	}
}
