package core

type DecodeFunc func(FetchedInstr) DecodedInstr

type DecodedInstr interface {
	Execute(*Registers)
}

type SetFlags uint8

const (
	ALWAYS SetFlags = iota
	NEVER
	NOT_IT // Only set condition codes if not in IT block
)

type InstrFields struct {
	setflags SetFlags
	Imm      uint32
	Rd       RegIndex
	Rm       RegIndex
	Rn       RegIndex
}
