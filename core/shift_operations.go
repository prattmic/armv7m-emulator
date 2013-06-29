package core

/* Perform LSL instruction, updating condition codes */
func LSL(regs *Registers, value uint32, shift_n uint8, S bool) uint32 {
	var result uint32
	var carry_out bool

	if shift_n == 0 {
		result, carry_out = value, regs.Apsr.C
	} else {
		result, carry_out = LSL_C(value, shift_n)
	}

	if S {
		regs.Apsr.N = (result & 0x80000000) != 0
		regs.Apsr.Z = (result) == 0
		regs.Apsr.C = carry_out
	}

	return result
}

/* Left shift value by a positive amount */
func LSL_C(value uint32, amount uint8) (uint32, bool) {
	extended := uint64(value)

	extended = extended << amount

	result := uint32(extended & 0xffffffff)
	carry_out := (extended & 0x100000000) != 0

	return result, carry_out
}
