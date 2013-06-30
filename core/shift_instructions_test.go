package core

import (
	"reflect"
	"testing"
)

func TestIdentifyLslImm(t *testing.T) {
	cases := []IdentifyCase{
		{instr: FetchedInstr16(0x0040), instr_valid: true},  // lsl r0, r0, #1
		{instr: FetchedInstr16(0x01e7), instr_valid: true},  // lsl r7, r4, #7
		{instr: FetchedInstr16(0x0000), instr_valid: false}, // lsl r0, r0, #0 -> mov r0, r0
		{instr: FetchedInstr16(0x0001), instr_valid: false}, // lsl r1, r0, #0 -> mov r1, r0
		{instr: FetchedInstr16(0x4080), instr_valid: false}, // lsl r0, r0, r0
		{instr: FetchedInstr16(0x40bf), instr_valid: false}, // lsl r7, r7, r7
		{instr: FetchedInstr16(0x09e7), instr_valid: false}, // lsr r7, r4, #7
		{instr: FetchedInstr16(0xffff), instr_valid: false},
	}

	test_identify(t, cases, reflect.TypeOf(LslImm{}))
}

func TestDecodeLslImm16(t *testing.T) {
	cases := []DecodeCase{
		// lsl r0, r0, #1
		{instr: FetchedInstr16(0x0040), decoded: LslImm{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT}},
		// lsl r7, r4, #7
		{instr: FetchedInstr16(0x01e7), decoded: LslImm{Rd: 7, Rm: 4, Rn: 0, Imm: 7, setflags: NOT_IT}},
		// lsl r0, r0, #0 -> mov r0, r0
		{instr: FetchedInstr16(0x0000), decoded: MovRegT2{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: ALWAYS}},
	}

	test_decode(t, cases, LslImm16)
}

func TestExecuteLslImm(t *testing.T) {
	cases := []ExecuteCase{
		// lsl r0, r0, #1
		{instr: LslImm{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}},
		// lsl r7, r4, #7
		{instr: LslImm{Rd: 7, Rm: 4, Rn: 0, Imm: 7, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{3, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{3, 0, 0, 0, 1, 0, 0, 0x80, 0, 0, 0, 0, 0}}},
		// lsl r0, r0, #1
		{instr: LslImm{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0xc0000000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0x80000000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{C: true, N: true}}},
		// lsl r0, r0, #1
		{instr: LslImm{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0x80000000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{C: true, Z: true}}},
	}

	test_execute(t, cases)
}

func TestIdentifyLslReg(t *testing.T) {
	cases := []IdentifyCase{
		{instr: FetchedInstr16(0x4080), instr_valid: true},  // lsl r0, r0, r0
		{instr: FetchedInstr16(0x40bf), instr_valid: true},  // lsl r7, r7, r7
		{instr: FetchedInstr16(0x0000), instr_valid: false}, // lsl r0, r0, #0
		{instr: FetchedInstr16(0x01e7), instr_valid: false}, // lsl r7, r4, #7
		{instr: FetchedInstr16(0x09e7), instr_valid: false}, // lsr r7, r4, #7
		{instr: FetchedInstr16(0xffff), instr_valid: false},
	}

	test_identify(t, cases, reflect.TypeOf(LslReg{}))
}

func TestDecodeLslReg16(t *testing.T) {
	cases := []DecodeCase{
		// lsl r0, r0, r0
		{instr: FetchedInstr16(0x4080), decoded: LslReg{Rd: 0, Rn: 0, Rm: 0, Imm: 0, setflags: NOT_IT}},
		// lsl r7, r7, r7
		{instr: FetchedInstr16(0x40bf), decoded: LslReg{Rd: 7, Rn: 7, Rm: 7, Imm: 0, setflags: NOT_IT}},
		// lsl r4, r4, r7
		{instr: FetchedInstr16(0x40bc), decoded: LslReg{Rd: 4, Rn: 4, Rm: 7, Imm: 0, setflags: NOT_IT}},
	}

	test_decode(t, cases, LslReg16)
}

func TestExecuteLslReg(t *testing.T) {
	cases := []ExecuteCase{
		// lsl r0, r0, r0
		{instr: LslReg{Rd: 0, Rn: 0, Rm: 0, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}},
		// lsl r7, r7, r7
		{instr: LslReg{Rd: 7, Rn: 7, Rm: 7, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 0x800, 0, 0, 0, 0, 0}}},
		// lsl r4, r4, r7
		{instr: LslReg{Rd: 4, Rn: 4, Rm: 7, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{1, 2, 3, 4, 0x500, 6, 7, 8, 0, 0, 0, 0, 0}}},
		// lsl r0, r0, r0
		// The bottom byte of the source register is the amount to shift by
		{instr: LslReg{Rd: 0, Rn: 0, Rm: 0, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0x100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0x100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}},
		// lsl r0, r0, r1
		{instr: LslReg{Rd: 0, Rn: 0, Rm: 1, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0xc0000000, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0x80000000, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{C: true, N: true}}},
		// lsl r0, r0, r1
		{instr: LslReg{Rd: 0, Rn: 0, Rm: 1, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0x80000000, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{C: true, Z: true}}},
	}

	test_execute(t, cases)
}

func TestIdentifyLsrImm(t *testing.T) {
	cases := []IdentifyCase{
		{instr: FetchedInstr16(0x0800), instr_valid: true},  // lsr r0, r0, #0
		{instr: FetchedInstr16(0x09e7), instr_valid: true},  // lsr r7, r4, #7
		{instr: FetchedInstr16(0x40c0), instr_valid: false}, // lsr r0, r0, r0
		{instr: FetchedInstr16(0x40d3), instr_valid: false}, // lsr r3, r3, r1
		{instr: FetchedInstr16(0x01e7), instr_valid: false}, // lsl r7, r4, #7
		{instr: FetchedInstr16(0x0000), instr_valid: false}, // lsl r0, r0, #0
		{instr: FetchedInstr16(0xffff), instr_valid: false},
	}

	test_identify(t, cases, reflect.TypeOf(LsrImm{}))
}

func TestDecodeLsrImm16(t *testing.T) {
	cases := []DecodeCase{
		// lsr r0, r0, #0
		{instr: FetchedInstr16(0x0800), decoded: LsrImm{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT}},
		// lsr r0, r0, #1
		{instr: FetchedInstr16(0x0840), decoded: LsrImm{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT}},
		// lsr r7, r4, #7
		{instr: FetchedInstr16(0x09e7), decoded: LsrImm{Rd: 7, Rm: 4, Rn: 0, Imm: 7, setflags: NOT_IT}},
	}

	test_decode(t, cases, LsrImm16)
}

func TestExecuteLsrImm(t *testing.T) {
	cases := []ExecuteCase{
		// lsr r0, r0, #0
		{instr: LsrImm{Rd: 0, Rm: 0, Rn: 0, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}},
		// lsr r0, r0, #1
		{instr: LsrImm{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}},
		// lsr r7, r4, #7
		{instr: LsrImm{Rd: 7, Rm: 4, Rn: 0, Imm: 7, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 0, 0, 0, 0x80, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{1, 0, 0, 0, 0x80, 0, 0, 1, 0, 0, 0, 0, 0}}},
		// lsr r0, r0, #1
		{instr: LsrImm{Rd: 0, Rm: 0, Rn: 0, Imm: 1, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{C: true, Z: true}}},
	}

	test_execute(t, cases)
}

func TestIdentifyLsrReg(t *testing.T) {
	cases := []IdentifyCase{
		{instr: FetchedInstr16(0x40c0), instr_valid: true},  // lsr r0, r0, r0
		{instr: FetchedInstr16(0x40d3), instr_valid: true},  // lsr r3, r3, r1
		{instr: FetchedInstr16(0x0800), instr_valid: false}, // lsr r0, r0, #0
		{instr: FetchedInstr16(0x09e7), instr_valid: false}, // lsr r7, r4, #7
		{instr: FetchedInstr16(0x01e7), instr_valid: false}, // lsl r7, r4, #7
		{instr: FetchedInstr16(0x0000), instr_valid: false}, // lsl r0, r0, #0
		{instr: FetchedInstr16(0xffff), instr_valid: false},
	}

	test_identify(t, cases, reflect.TypeOf(LsrReg{}))
}

func TestDecodeLsrReg16(t *testing.T) {
	cases := []DecodeCase{
		// lsr r0, r0, r0
		{instr: FetchedInstr16(0x40c0), decoded: LsrReg{Rd: 0, Rn: 0, Rm: 0, Imm: 0, setflags: NOT_IT}},
		// lsr r7, r7, r7
		{instr: FetchedInstr16(0x40ff), decoded: LsrReg{Rd: 7, Rn: 7, Rm: 7, Imm: 0, setflags: NOT_IT}},
		// lsr r4, r4, r7
		{instr: FetchedInstr16(0x40fc), decoded: LsrReg{Rd: 4, Rn: 4, Rm: 7, Imm: 0, setflags: NOT_IT}},
	}

	test_decode(t, cases, LsrReg16)
}

func TestExecuteLsrReg(t *testing.T) {
	cases := []ExecuteCase{
		// lsr r0, r0, r0
		{instr: LsrReg{Rd: 0, Rn: 0, Rm: 0, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{C: true, Z: true}}},
		// lsl r4, r4, r7
		{instr: LsrReg{Rd: 4, Rn: 4, Rm: 7, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 2, 3, 4, 5, 6, 7, 2, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{1, 2, 3, 4, 1, 6, 7, 2, 0, 0, 0, 0, 0}}},
		// lsl r0, r0, r0
		// The bottom byte of the source register is the amount to shift by
		{instr: LsrReg{Rd: 0, Rn: 0, Rm: 0, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0x100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0x100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}},
		// lsr r0, r0, r1
		{instr: LsrReg{Rd: 0, Rn: 0, Rm: 0, Imm: 0, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{Z: true}}},
	}

	test_execute(t, cases)
}

func TestIdentifyAsrImm(t *testing.T) {
	cases := []IdentifyCase{
		{instr: FetchedInstr16(0x1000), instr_valid: true},  // asr r0, r0, #32 (imm of 0 becomes imm of 32)
		{instr: FetchedInstr16(0x1001), instr_valid: true},  // asr r1, r0, #32
		{instr: FetchedInstr16(0x111a), instr_valid: true},  // asr r2, r1, #2
		{instr: FetchedInstr16(0x09e7), instr_valid: false}, // lsr r7, r4, #7
		{instr: FetchedInstr16(0x40c0), instr_valid: false}, // lsr r0, r0, r0
		{instr: FetchedInstr16(0x40d3), instr_valid: false}, // lsr r3, r3, r1
		{instr: FetchedInstr16(0x01e7), instr_valid: false}, // lsl r7, r4, #7
		{instr: FetchedInstr16(0x0000), instr_valid: false}, // lsl r0, r0, #0
		{instr: FetchedInstr16(0xffff), instr_valid: false},
	}

	test_identify(t, cases, reflect.TypeOf(AsrImm{}))
}

func TestDecodeAsrImm16(t *testing.T) {
	cases := []DecodeCase{
		// asr r0, r0, #32
		{instr: FetchedInstr16(0x1000), decoded: AsrImm{Rd: 0, Rm: 0, Rn: 0, Imm: 32, setflags: NOT_IT}},
		// asr r1, r0, #32
		{instr: FetchedInstr16(0x1001), decoded: AsrImm{Rd: 1, Rm: 0, Rn: 0, Imm: 32, setflags: NOT_IT}},
		// asr r2, r3, #4
		{instr: FetchedInstr16(0x111a), decoded: AsrImm{Rd: 2, Rm: 3, Rn: 0, Imm: 4, setflags: NOT_IT}},
	}

	test_decode(t, cases, AsrImm16)
}

func TestExecuteAsrImm(t *testing.T) {
	cases := []ExecuteCase{
		// asr r0, r0, #32
		{instr: AsrImm{Rd: 0, Rm: 0, Rn: 0, Imm: 32, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{Z: true}}},
		// asr r0, r0, #32
		{instr: AsrImm{Rd: 0, Rm: 0, Rn: 0, Imm: 32, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0x80000000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0xffffffff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{N: true, C: true}}},
		// asr r2, r1, #2
		{instr: AsrImm{Rd: 2, Rm: 1, Rn: 0, Imm: 2, setflags: NOT_IT},
			regs:     Registers{r: GeneralRegs{0, 0xffffffff - 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: Registers{r: GeneralRegs{0, 0xffffffff - 8, 0xffffffff - 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Apsr: Apsr{N: true, C: true}}},
	}

	test_execute(t, cases)
}
