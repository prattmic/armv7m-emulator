package core

/* LSL - Logical Shift Left
 * ARM ARM A7.7.67 - A7.7.68 */
type Lsl InstrFields

func LslImm16(instr FetchedInstr) DecodedInstr {
    raw_instr := instr.Uint32()

    Rd := uint8(raw_instr & 0x7)
    Rm := uint8((raw_instr >> 3) & 0x7)
    Imm := (raw_instr >> 6) & 0x1f

    return Lsl{Rd: Rd, Rm: Rm, Imm: Imm, S: true, Rn: 0}
}

func LslReg16(instr FetchedInstr) DecodedInstr {
    raw_instr := instr.Uint32()

    Rdn := uint8(raw_instr & 0x7)
    Rm := uint8((raw_instr >> 3) & 0x7)

    return Lsl{Rd: Rdn, Rn: Rdn, Rm: Rm, Imm: 0, S: true}
}

func (instr Lsl) Execute(regs *Registers) {
    regs.R[instr.Rd] = regs.R[instr.Rm] << instr.Imm
}

/* LSR - Logical Shift Right
 * ARM ARM A7.7.69 - A7.7.69 */
type Lsr InstrFields

func LsrImm16(instr FetchedInstr) DecodedInstr {
    raw_instr := instr.Uint32()

    Rd := uint8(raw_instr & 0x7)
    Rm := uint8((raw_instr >> 3) & 0x7)
    Imm := (raw_instr >> 6) & 0x1f

    return Lsr{Rd: Rd, Rm: Rm, Imm: Imm, S: true, Rn: 0}
}

func LsrReg16(instr FetchedInstr) DecodedInstr {
    raw_instr := instr.Uint32()

    Rdn := uint8(raw_instr & 0x7)
    Rm := uint8((raw_instr >> 3) & 0x7)

    return Lsr{Rd: Rdn, Rn: Rdn, Rm: Rm, Imm: 0, S: true}
}

func (instr Lsr) Execute(regs *Registers) {
    regs.R[instr.Rd] = regs.R[instr.Rm] >> instr.Imm
}
