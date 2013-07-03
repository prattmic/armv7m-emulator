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

func (setflags SetFlags) String() string {
	if setflags == ALWAYS || setflags == NOT_IT {
		return "s"
	}

	return ""
}

func (setflags SetFlags) ShouldSetFlags(regs Registers) bool {
	if setflags == ALWAYS || (setflags == NOT_IT && !regs.InITBlock()) {
		return true
	}
	return false
}

type InstrFields struct {
	setflags SetFlags
	Imm      uint32
	Rd       RegIndex
	Rm       RegIndex
	Rn       RegIndex
}
