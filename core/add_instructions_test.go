package core

import (
	"reflect"
	"testing"
)

func TestIdentifyAddRegT1(t *testing.T) {
	cases := []IdentifyCase{
		{instr: FetchedInstr16(0x1800), instr_valid: true},  // adds r0, r0, r0
		{instr: FetchedInstr16(0x19ff), instr_valid: true},  // adds r7, r7, r7
		{instr: FetchedInstr16(0x18d1), instr_valid: true},  // adds r1, r2, r3
		{instr: FetchedInstr16(0x0000), instr_valid: false}, // mov r0, r0
		{instr: FetchedInstr16(0x001f), instr_valid: false}, // mov r7, r3
		{instr: FetchedInstr16(0x2000), instr_valid: false}, // mov r0, #0
		{instr: FetchedInstr16(0x2745), instr_valid: false}, // mov r7, #0x45
		{instr: FetchedInstr16(0x4080), instr_valid: false}, // lsl r0, r0, r0
		{instr: FetchedInstr16(0xffff), instr_valid: false},
	}

	test_identify(t, cases, reflect.TypeOf(AddRegT1{}))
}

func TestDecodeAddReg16T1(t *testing.T) {
	cases := []DecodeCase{
		// adds r0, r0, r0
		{instr: FetchedInstr16(0x1800), decoded: AddRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT}},
		// adds r7, r7, r7
		{instr: FetchedInstr16(0x19ff), decoded: AddRegT1{Rd: 7, Rm: 7, Rn: 7, Imm: 0, setflags: NOT_IT}},
		// adds r1, r2, r3
		{instr: FetchedInstr16(0x18d1), decoded: AddRegT1{Rd: 1, Rm: 3, Rn: 2, Imm: 0, setflags: NOT_IT}},
	}

	test_decode(t, cases, AddReg16T1)
}

func TestExecuteAddRegT1(t *testing.T) {
	cases := []ExecuteCase{
		// adds r0, r0, r0
		{instr: AddRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
			expected: Registers{r: GeneralRegs{0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true}}},
		// adds r0, r0, r0
		{instr: AddRegT1{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: true}},
			expected: Registers{r: GeneralRegs{2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, Apsr: Apsr{Z: false}}},
		// adds r1, r2, r3
		{instr: AddRegT1{Rd: 1, Rm: 3, Rn: 2, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
			expected: Registers{r: GeneralRegs{1, 7, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}}},
		// adds r0, r1, r2
		{instr: AddRegT1{Rd: 0, Rm: 2, Rn: 1, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0, 0x7fffffff, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
			expected: Registers{r: GeneralRegs{0x80000000, 0x7fffffff, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{V: true, N: true}}},
		// adds r0, r1, r1
		{instr: AddRegT1{Rd: 0, Rm: 1, Rn: 1, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0, 0x80000000, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
			expected: Registers{r: GeneralRegs{0, 0x80000000, 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Apsr: Apsr{Z: true, C: true, V: true}}},
	}

	test_execute(t, cases)
}
