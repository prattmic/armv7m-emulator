package core

import (
	"bytes"
	"fmt"
)

type GeneralRegs [13]uint32
type RegIndex uint8

type SPRegs [2]uint32
type SPType uint8

const (
	MSP SPType = 0
	PSP SPType = 1
)

type Apsr struct {
	N  bool  // Negative
	Z  bool  // Zero
	C  bool  // Carry
	V  bool  // Overflow
	Q  bool  // Saturation
	GE uint8 // Greater than or equal flags
}

type Ipsr struct {
	ExcpNum uint16
}

type Epsr struct {
	T   bool   // Thumb bit
	ICI uint16 // Interrupt-continue
	IT  uint16 // IT block flags
}

type Mode uint8

const (
	MODE_THREAD Mode = iota
	MODE_HANDLER
)

type Control struct {
	Npriv bool   // NOT privilege
	Spsel SPType // SP select
	Fpca  bool   // FP extension enable
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
	SP RegIndex = 13
	LR RegIndex = 14
	PC RegIndex = 15
)

func (regs Registers) R(i RegIndex) uint32 {
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

func (regs *Registers) SetR(i RegIndex, value uint32) {
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

func (regs Registers) Pretty() string {
	var b bytes.Buffer
	var i RegIndex

	for i = 0; i <= 12; i++ {
		if i != 0 {
			if (i % 4) == 0 {
				fmt.Fprintf(&b, "\n")
			} else {
				fmt.Fprintf(&b, "\t")
			}
		}
		fmt.Fprintf(&b, "R%-2d = %#x", i, regs.R(i))
	}

	fmt.Fprintf(&b, "\tSP (R13) = %#x", regs.R(SP))
	fmt.Fprintf(&b, "\tLR (R14) = %#x", regs.R(LR))
	fmt.Fprintf(&b, "\tPC (R15) = %#x\n", regs.R(PC))

	fmt.Fprintf(&b, "MSP = %#x\tPSP = %#x\n", regs.Msp(), regs.Psp())

	fmt.Fprintf(&b, "BASEPRI = %d\tPRIMASK = %d\tFAULTMASK = %d\n", regs.Basepri,
		booltou(regs.Primask), booltou(regs.Faultmask))

	fmt.Fprintf(&b, "CONTROL: nPRIV = %d SPSEL = %d FPCA = %d\n", booltou(regs.Control.Npriv),
		uint8(regs.Control.Spsel), booltou(regs.Control.Fpca))

	fmt.Fprintf(&b, "APSR: N = %d Z = %d C = %d V = %d Q = %d GE = %d\n",
		booltou(regs.Apsr.N), booltou(regs.Apsr.Z), booltou(regs.Apsr.C),
		booltou(regs.Apsr.V), booltou(regs.Apsr.Q), regs.Apsr.GE)

	fmt.Fprintf(&b, "EPSR: T = %d ICI = %#x IT = %#x\n", booltou(regs.Epsr.T),
		regs.Epsr.ICI, regs.Epsr.IT)

	fmt.Fprintf(&b, "IPSR: EXCPNUM = %d\n", regs.Ipsr.ExcpNum)

	return b.String()
}

func (regs Registers) Print() {
	fmt.Print(regs.Pretty())
}

func (i RegIndex) String() string {
	switch i {
	default:
		return fmt.Sprintf("r%d", i)
	case SP:
		return "sp"
	case LR:
		return "lr"
	case PC:
		return "pc"
	}
}
