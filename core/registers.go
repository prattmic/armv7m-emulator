package core

import "fmt"

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
	T   bool
	ICI uint16
	IT  uint16
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

func (regs Registers) Print() {
	for i := 0; i < 16; i++ {
		if i != 0 {
			if (i % 4) == 0 {
				fmt.Printf("\n")
			} else {
				fmt.Printf("\t")
			}
		}
		fmt.Printf("R%-2d = 0x%x", i, regs.R[i])
	}

	fmt.Printf("\nAPSR: N = %d Z = %d C = %d V = %d Q = %d GE = %d\n",
		booltoi(regs.Apsr.N), booltoi(regs.Apsr.Z), booltoi(regs.Apsr.C),
		booltoi(regs.Apsr.V), booltoi(regs.Apsr.Q), regs.Apsr.GE)

	fmt.Printf("EPSR: T = %d ICI = 0x%x IT = 0x%x\n", booltoi(regs.Epsr.T),
		regs.Epsr.ICI, regs.Epsr.IT)

	fmt.Printf("IPSR: %d\n", regs.Ipsr.ExcpNum)
}

func booltoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
