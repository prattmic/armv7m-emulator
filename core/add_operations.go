package core

/* Perform actual addition operation, determining carry out and overflow */
func AddWithCarry(x uint32, y uint32, carry_in uint8) (uint32, uint8, uint8) {
	var unsigned_sum uint64 = uint64(x) + uint64(y) + uint64(carry_in)
	var signed_sum int64 = int64(int32(x)) + int64(int32(y)) + int64(carry_in)
	var result uint64 = unsigned_sum & 0xffffffff

	var carry_out uint8
	var overflow uint8

	if result == unsigned_sum {
		carry_out = 0
	} else {
		carry_out = 1
	}

	if int64(int32(result)) == signed_sum {
		overflow = 0
	} else {
		overflow = 1
	}

	return uint32(result), carry_out, overflow
}

/* Perform addition instruction, with shift, updating condition codes */
func AddRegister(regs *Registers, instr InstrFields, shift Shift) {
	shifted, _ := shift.Evaluate(regs.R(instr.Rm))
	result, carry, overflow := AddWithCarry(regs.R(instr.Rn), shifted, 0)

	if instr.Rd == PC {
		regs.ALUWritePC(result)
	} else {
		regs.SetR(instr.Rd, result)
		if instr.setflags.ShouldSetFlags(*regs) {
			regs.Apsr.N = (result & 0x80000000) != 0
			regs.Apsr.Z = (result) == 0
			regs.Apsr.C = utobool(carry)
			regs.Apsr.V = utobool(overflow)
		}
	}
}
