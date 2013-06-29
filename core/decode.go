package core

import (
    "fmt"
    "errors"
)

/* If bits [15:11] of the instruction being decoded are any of
 * the following WORD_INSTR, then the instruction is the first
 * halfword of a 32-bit instruction.
 * ARMv7-M ARM A5.1 */
const (
    WORD_INSTR_MASK = 0xf800
    WORD_INSTR1 = 0xe800
    WORD_INSTR2 = 0xf000
    WORD_INSTR3 = 0xf800
)

var IncompleteInstruction = errors.New("Only first halfword of word instruction decoded.")
var UndefinedInstruction = errors.New("Instruction not defined.")

type FetchedInstr interface {
    Decode() (DecodedInstr, error)
    String() string
    Uint32() uint32
}

type FetchedInstr16 uint16
type FetchedInstr32 uint32

func (instr FetchedInstr16) Decode() (DecodedInstr, error) {
    raw_instr := uint16(instr)

    /* Check if this is the beginning of a 32-bit instruction */
    switch (raw_instr & WORD_INSTR_MASK) {
    case WORD_INSTR1, WORD_INSTR2, WORD_INSTR3:
        return nil, IncompleteInstruction
    }

    /* Check for a matching opcode */
    for opcode, decode := range InstrOpcodes16 {
        if opcode.Match(instr) {
            /* Instruction identified, now decode it */
            return decode(instr), nil
        }
    }

    return nil, UndefinedInstruction
}

func (instr FetchedInstr16) Uint32() uint32 {
    return uint32(instr)
}

func (instr FetchedInstr16) String() string {
    return fmt.Sprintf("0x%.4x", uint16(instr))
}

/* Extend upper halfword of instruction with lower halfword to make 32-bit instruction */
func (upper FetchedInstr16) Extend(lower FetchedInstr16) FetchedInstr32 {
    return FetchedInstr32((uint32(upper) << 16) | uint32(lower))
}

func (instr FetchedInstr32) Decode() (DecodedInstr, error) {
    /* Check for a matching opcode */
    for opcode, decode := range InstrOpcodes32 {
        if opcode.Match(instr) {
            return decode(instr), nil
        }
    }

    return nil, UndefinedInstruction
}

func (instr FetchedInstr32) Uint32() uint32 {
    return uint32(instr)
}

func (instr FetchedInstr32) String() string {
    return fmt.Sprintf("0x%.8x", uint32(instr))
}
