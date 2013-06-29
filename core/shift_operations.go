package core

/* Left shift value by a positive amount */
func LSL_C(value uint32, amount uint8) (uint32, bool) {
	extended := uint64(value)

	extended = extended << amount

	result := uint32(extended & 0xffffffff)
	carry_out := (extended & 0x100000000) != 0

	return result, carry_out
}
