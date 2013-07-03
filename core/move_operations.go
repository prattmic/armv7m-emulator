package core

/* Move value into destination register, updating condition codes */
func MoveValue(regs *Registers, dest RegIndex, value uint32, setflags SetFlags, carry bool) {
	regs.SetR(dest, value)

	if setflags.ShouldSetFlags(*regs) {
		regs.Apsr.N = (value & 0x80000000) != 0
		regs.Apsr.Z = value == 0
		regs.Apsr.C = carry
	}
}

func MoveRegister(regs *Registers, dest RegIndex, source RegIndex, setflags SetFlags, carry bool) {
	value := regs.R(source)

	if dest == PC {
		regs.ALUWritePC(value)
	} else {
		MoveValue(regs, dest, value, setflags, carry)
	}
}
