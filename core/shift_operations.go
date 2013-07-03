package core

/* Generic shifting function
 *
 * @param value uint32		Value to shift
 * @param amount uint8		Amount to shift by
 *
 * @return result uint32	Result of shift
 * @return carry bool		Carry out of shift
 */
type ShiftFunc func(uint32, uint8) (uint32, bool)

type Shift struct {
	function ShiftFunc
	amount   uint8
}

func (shift Shift) Evaluate(input uint32) (uint32, bool) {
	return shift.function(input, shift.amount)
}

/* Perform shift operation, updating condition codes */
func ShiftOp(regs *Registers, value uint32, shift_n uint8, setflags SetFlags, do_shift ShiftFunc) uint32 {
	var result uint32
	var carry_out bool

	if shift_n == 0 {
		result, carry_out = value, regs.Apsr.C
	} else {
		result, carry_out = do_shift(value, shift_n)
	}

	if setflags.ShouldSetFlags(*regs) {
		regs.Apsr.N = (result & 0x80000000) != 0
		regs.Apsr.Z = (result) == 0
		regs.Apsr.C = carry_out
	}

	return result
}

/* Perform LSL instruction, updating condition codes */
func LSL(regs *Registers, value uint32, shift_n uint8, setflags SetFlags) uint32 {
	return ShiftOp(regs, value, shift_n, setflags, LSL_C)
}

/* Left shift value by a positive amount */
func LSL_C(value uint32, amount uint8) (uint32, bool) {
	extended := uint64(value)

	extended = extended << amount

	result := uint32(extended & 0xffffffff)
	carry_out := (extended & 0x100000000) != 0

	return result, carry_out
}

/* Perform LSR instruction, updating condition codes */
func LSR(regs *Registers, value uint32, shift_n uint8, setflags SetFlags) uint32 {
	return ShiftOp(regs, value, shift_n, setflags, LSR_C)
}

/* Right shift value by a positive amount */
func LSR_C(value uint32, amount uint8) (uint32, bool) {
	/* The last bit to be carried out determines the carry */
	carry_out := (value & (1 << (amount - 1))) != 0

	result := value >> amount

	return result, carry_out
}

/* Perform ASR instruction, updating condition codes */
func ASR(regs *Registers, value uint32, shift_n uint8, setflags SetFlags) uint32 {
	return ShiftOp(regs, value, shift_n, setflags, ASR_C)
}

/* Right shift value by a positive amount, copying the leftmost bit */
func ASR_C(value uint32, amount uint8) (uint32, bool) {
	/* The last bit to be carried out determines the carry */
	carry_out := (value & (1 << (amount - 1))) != 0

	extended := int32(value)

	result := extended >> amount

	return uint32(result), carry_out
}
