package core

type GeneralRegs [16]uint32

type Apsr struct {
	N  bool
	Z  bool
	C  bool
	V  bool
	Q  bool
	GE uint8
}

type Ipsr struct {
	ExcpNum uint16
}

type Epsr struct {
	T      bool
	ICI_IT uint16
}

type Registers struct {
	R    GeneralRegs
	Apsr Apsr
	Ipsr Ipsr
	Epsr Epsr
}

/* Special registers in r13-15 */
const (
	LR = 13
	SP = 14
	PC = 15
)

func (regs Registers) Lr() uint32 {
	return regs.R[LR]
}

func (regs Registers) Sp() uint32 {
	return regs.R[SP]
}

func (regs Registers) Pc() uint32 {
	return regs.R[PC]
}
