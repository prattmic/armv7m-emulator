package core

import (
	"reflect"
	"testing"
)

func TestIdentifyMovImm(t *testing.T) {
	cases := []IdentifyCase{
		{instr: FetchedInstr16(0x2000), instr_valid: true},  // mov r0, #0
		{instr: FetchedInstr16(0x2745), instr_valid: true},  // mov r7, #0x45
		{instr: FetchedInstr16(0x4080), instr_valid: false}, // lsl r0, r0, r0
		{instr: FetchedInstr16(0x40bf), instr_valid: false}, // lsl r7, r7, r7
		{instr: FetchedInstr16(0x09e7), instr_valid: false}, // lsr r7, r4, #7
		{instr: FetchedInstr16(0xffff), instr_valid: false},
	}

	test_identify(t, cases, reflect.TypeOf(MovImm{}))
}

func TestDecodeMovImm16(t *testing.T) {
	cases := []DecodeCase{
		// mov r0, #0
		{instr: FetchedInstr16(0x2000), decoded: MovImm{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT}},
		// mov r7, #0x45
		{instr: FetchedInstr16(0x2745), decoded: MovImm{Rd: 7, Rm: 0, Rn: 0, Imm: 0x45, setflags: NOT_IT}},
	}

	test_decode(t, cases, MovImm16)
}

func TestExecuteMovImm(t *testing.T) {
	cases := []ExecuteCase{
		// mov r0, #0
		{instr: MovImm{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT},
			regs:     Registers{R: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{R: GeneralRegs{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{Z: true}}},
		// mov r7, #0x45
		{instr: MovImm{Rd: 7, Rm: 0, Rn: 0, Imm: 0x45, setflags: NOT_IT},
			regs:     Registers{R: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{C: true}},
			expected: Registers{R: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0x45, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{C: true}}},
	}

	test_execute(t, cases)
}

func TestIdentifyMovRegT1(t *testing.T) {
	cases := []IdentifyCase{
		{instr: FetchedInstr16(0x4600), instr_valid: true},  // mov r0, r0
		{instr: FetchedInstr16(0x46e7), instr_valid: true},  // mov pc, r12
		{instr: FetchedInstr16(0x2000), instr_valid: false}, // mov r0, #0
		{instr: FetchedInstr16(0x2745), instr_valid: false}, // mov r7, #0x45
		{instr: FetchedInstr16(0x4080), instr_valid: false}, // lsl r0, r0, r0
		{instr: FetchedInstr16(0xffff), instr_valid: false},
	}

	test_identify(t, cases, reflect.TypeOf(MovRegT1{}))
}

func TestDecodeMovReg16T1(t *testing.T) {
	cases := []DecodeCase{
		// mov r0, r0
		{instr: FetchedInstr16(0x4600), decoded: MovRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NEVER}},
		// mov pc, r12
		{instr: FetchedInstr16(0x46e7), decoded: MovRegT1{Rd: 15, Rm: 12, Rn: 0, Imm: 0, setflags: NEVER}},
	}

	test_decode(t, cases, MovReg16T1)
}

func TestExecuteMovRegT1(t *testing.T) {
	cases := []ExecuteCase{
		// mov r0, r0
		{instr: MovRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NEVER},
			regs:     Registers{R: GeneralRegs{0xDEAD, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{Z: true}},
			expected: Registers{R: GeneralRegs{0xDEAD, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{Z: true}}},
		// mov r1, r0
		{instr: MovRegT1{Rd: 1, Rm: 0, Rn: 0, Imm: 0, setflags: NEVER},
			regs:     Registers{R: GeneralRegs{0xF00D, 0xDEAD, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{Z: true}},
			expected: Registers{R: GeneralRegs{0xF00D, 0xF00D, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{Z: true}}},
		// mov pc, r12
		{instr: MovRegT1{Rd: 15, Rm: 12, Rn: 0, Imm: 0, setflags: NEVER},
			regs:     Registers{R: GeneralRegs{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0xCAFE, 13, 14, 0xDEAD}, Apsr: Apsr{C: true}},
			expected: Registers{R: GeneralRegs{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0xCAFE, 13, 14, 0xCAFE}, Apsr: Apsr{C: true}}},
		// mov pc, r12
		{instr: MovRegT1{Rd: 15, Rm: 12, Rn: 0, Imm: 0, setflags: NEVER},
			regs:     Registers{R: GeneralRegs{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0xCAFF, 13, 14, 0xDEAD}, Apsr: Apsr{C: true}},
			expected: Registers{R: GeneralRegs{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0xCAFF, 13, 14, 0xCAFE}, Apsr: Apsr{C: true}}},
	}

	test_execute(t, cases)
}
