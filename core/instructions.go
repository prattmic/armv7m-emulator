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
