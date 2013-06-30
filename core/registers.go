package core

import "fmt"

type GeneralRegs [13]uint32

type SPRegs [2]uint32
type SPType uint8

const (
	MSP SPType = 0
	PSP SPType = 1
)

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

type Mode uint8

const (
	MODE_THREAD Mode = iota
	MODE_HANDLER
)

type Control struct {
	Npriv bool
	Spsel SPType
	Fpca  bool
}

type Registers struct {
	r         GeneralRegs
	sp        SPRegs
	lr        uint32
	pc        uint32
	Apsr      Apsr
	Ipsr      Ipsr
	Epsr      Epsr
	Mode      Mode
	Primask   bool
	Faultmask bool
	Basepri   uint8
	Control   Control
}

/* Special registers in r13-15 */
const (
	SP = 13
	LR = 14
	PC = 15
)

func (regs Registers) R(i uint8) uint32 {
	switch i {
	default:
		return regs.r[i]
	case SP:
		sp := regs.LookupSP()
		return regs.sp[sp]
	case LR:
		return regs.lr
	case PC:
		return regs.pc
	}
}

func (regs *Registers) SetR(i uint8, value uint32) {
	switch i {
	default:
		regs.r[i] = value
	case SP:
		sp := regs.LookupSP()
		regs.sp[sp] = value
	case LR:
		regs.lr = value
	case PC:
		regs.pc = value
	}
}

func (regs Registers) LookupSP() SPType {
	if regs.Control.Spsel == PSP && regs.Mode == MODE_THREAD {
		return PSP
	}

	return MSP
}

func (regs Registers) Lr() uint32 {
	return regs.R(LR)
}

func (regs Registers) Sp() uint32 {
	return regs.R(SP)
}

func (regs Registers) Msp() uint32 {
	return regs.sp[MSP]
}

func (regs Registers) Psp() uint32 {
	return regs.sp[PSP]
}

func (regs Registers) Pc() uint32 {
	return regs.R(PC)
}

func (regs Registers) InITBlock() bool {
	return (regs.Epsr.IT & 0xf) != 0
}

func (regs Registers) LastInITBlock() bool {
	return (regs.Epsr.IT & 0xf) == 0x8
}

func (regs *Registers) BranchTo(addr uint32) {
	regs.SetR(PC, addr)
}

func (regs *Registers) BranchWritePC(addr uint32) {
	addr &^= 0x1 // Clear the thumb bit

	regs.BranchTo(addr)
}

func (regs *Registers) ALUWritePC(addr uint32) {
	regs.BranchWritePC(addr)
}

func (regs Registers) Print() {
	var i uint8

	for i = 0; i <= 12; i++ {
		if i != 0 {
			if (i % 4) == 0 {
				fmt.Printf("\n")
			} else {
				fmt.Printf("\t")
			}
		}
		fmt.Printf("R%-2d = %#x", i, regs.R(i))
	}

	fmt.Printf("\tSP (R13) = %#x", regs.R(SP))
	fmt.Printf("\tLR (R14) = %#x", regs.R(LR))
	fmt.Printf("\tPC (R15) = %#x\n", regs.R(PC))

	fmt.Printf("MSP = %#x\tPSP = %#x\n", regs.Msp(), regs.Psp())

	fmt.Printf("BASEPRI = %d\tPRIMASK = %d\tFAULTMASK = %d\n", regs.Basepri,
		booltoi(regs.Primask), booltoi(regs.Faultmask))

	fmt.Printf("CONTROL: nPRIV = %d SPSEL = %d FPCA = %d\n", booltoi(regs.Control.Npriv),
		uint8(regs.Control.Spsel), booltoi(regs.Control.Fpca))

	fmt.Printf("APSR: N = %d Z = %d C = %d V = %d Q = %d GE = %d\n",
		booltoi(regs.Apsr.N), booltoi(regs.Apsr.Z), booltoi(regs.Apsr.C),
		booltoi(regs.Apsr.V), booltoi(regs.Apsr.Q), regs.Apsr.GE)

	fmt.Printf("EPSR: T = %d ICI = %#x IT = %#x\n", booltoi(regs.Epsr.T),
		regs.Epsr.ICI, regs.Epsr.IT)

	fmt.Printf("IPSR: EXCPNUM = %d\n", regs.Ipsr.ExcpNum)
}

func booltoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
