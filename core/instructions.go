package core

type DecodeFunc func(FetchedInstr) DecodedInstr

type DecodedInstr interface {
    Execute(*Registers)
}

type InstrFields struct {
    S bool
    Imm uint32
    Rd uint8
    Rm uint8
    Rn uint8
}

type Lsl InstrFields
type Lsr InstrFields

func LslImm16(instr FetchedInstr) DecodedInstr {
    raw_instr := instr.Uint32()

    Rd := uint8(raw_instr & 0x7)
    Rm := uint8((raw_instr >> 3) & 0x7)
    Imm := (raw_instr >> 6) & 0x1f

    return Lsl{Rd: Rd, Rm: Rm, Imm: Imm, S: true, Rn: 0}
}

func LsrImm16(instr FetchedInstr) DecodedInstr {
    raw_instr := instr.Uint32()

    Rd := uint8(raw_instr & 0x7)
    Rm := uint8((raw_instr >> 3) & 0x7)
    Imm := (raw_instr >> 6) & 0x1f

    return Lsr{Rd: Rd, Rm: Rm, Imm: Imm, S: true, Rn: 0}
}

func (instr Lsl) Execute(regs *Registers) {
    regs.R[instr.Rd] = regs.R[instr.Rm] << instr.Imm
}

func (instr Lsr) Execute(regs *Registers) {
    regs.R[instr.Rd] = regs.R[instr.Rm] >> instr.Imm
}
