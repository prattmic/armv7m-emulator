package core

import (
	"reflect"
	"testing"
)

func TestLsl(t *testing.T) {
	cases := []IdentifyCase{
		{FetchedInstr16(0x0000), true},
		{FetchedInstr16(0x01e7), true},
		{FetchedInstr16(0x09e7), false},
		{FetchedInstr16(0xffff), false},
	}

	instr_type := reflect.TypeOf(Lsl{})

	for _, test := range cases {
		instr, err := test.instr.Decode()
		verify_identify(t, test, instr_type, instr, err)
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
		verify_decode(t, test, actual)
	}
}

func TestLsr(t *testing.T) {
	cases := []IdentifyCase{
		{FetchedInstr16(0x0800), true},
		{FetchedInstr16(0x09e7), true},
		{FetchedInstr16(0x01e7), false},
		{FetchedInstr16(0x0000), false},
		{FetchedInstr16(0xffff), false},
	}

	instr_type := reflect.TypeOf(Lsr{})

	for _, test := range cases {
		instr, err := test.instr.Decode()
		verify_identify(t, test, instr_type, instr, err)
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
		verify_decode(t, test, actual)
	}
}
