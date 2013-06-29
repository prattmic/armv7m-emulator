package core

type GeneralRegs [16]uint32

type Registers struct {
    R GeneralRegs
    Psr uint32
}

/* Special registers in r13-15 */
const (
    LR = 13
    SP = 14
    PC = 15
)

func (regs *Registers) Lr() uint32 {
    return regs.R[LR]
}

func (regs *Registers) Sp() uint32 {
    return regs.R[SP]
}

func (regs *Registers) Pc() uint32 {
    return regs.R[PC]
}
