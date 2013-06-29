package core

type UnpredictableInstr InstrFields

// Case to execute in the event of UNPREDICTABLE instruction behavior
func (instr UnpredictableInstr) Execute(regs *Registers) {
	// Do nothing, for now
	return
}

type UndefinedInstr InstrFields

// Placeholder UNDEFINED instruction
func (instr UndefinedInstr) Execute(regs *Registers) {
	// Do nothing, for now
	return
}
