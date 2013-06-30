package core

type Opcode struct {
	mask  uint32
	value uint32
}

func (op *Opcode) Match(instr FetchedInstr) bool {
	raw_instr := instr.Uint32()

	return (raw_instr & op.mask) == op.value
}

var InstrOpcodes16 = map[Opcode]DecodeFunc{
	Opcode{mask: 0xf800, value: 0x0000}: LslImm16,
	Opcode{mask: 0xffc0, value: 0x4080}: LslReg16,
	Opcode{mask: 0xf800, value: 0x0800}: LsrImm16,
	Opcode{mask: 0xffc0, value: 0x40c0}: LsrReg16,
	Opcode{mask: 0xf800, value: 0x2000}: MovImm16,
	Opcode{mask: 0xff00, value: 0x4600}: MovReg16T1,
}

var InstrOpcodes32 = map[Opcode]DecodeFunc{}
