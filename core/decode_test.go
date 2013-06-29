package core

import "testing"

type DecodeCase struct {
	instr   FetchedInstr
	decoded DecodedInstr
}

func verify(t *testing.T, instr FetchedInstr, expected DecodedInstr, actual DecodedInstr) {
	if actual != expected {
		t.Errorf("instr: %v, expected %#v, got %#v", instr, expected, actual)
	}
}

func TestLslImm16(t *testing.T) {
	cases := []DecodeCase{
		{FetchedInstr16(0x0000), Lsl{Rd: 0, Rm: 0, Imm: 0, S: true, Rn: 0}},
		{FetchedInstr16(0x0040), Lsl{Rd: 0, Rm: 0, Imm: 1, S: true, Rn: 0}},
		{FetchedInstr16(0x01e7), Lsl{Rd: 7, Rm: 4, Imm: 7, S: true, Rn: 0}},
	}

	for _, test := range cases {
		actual := LslImm16(test.instr)
		verify(t, test.instr, test.decoded, actual)
	}
}

func TestLsrImm16(t *testing.T) {
	cases := []DecodeCase{
		{FetchedInstr16(0x0800), Lsr{Rd: 0, Rm: 0, Imm: 0, S: true, Rn: 0}},
		{FetchedInstr16(0x0840), Lsr{Rd: 0, Rm: 0, Imm: 1, S: true, Rn: 0}},
		{FetchedInstr16(0x09e7), Lsr{Rd: 7, Rm: 4, Imm: 7, S: true, Rn: 0}},
	}

	for _, test := range cases {
		actual := LsrImm16(test.instr)
		verify(t, test.instr, test.decoded, actual)
	}
}
