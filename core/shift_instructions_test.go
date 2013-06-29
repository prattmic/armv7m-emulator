package core

import (
	"reflect"
	"testing"
)

func TestIdentifyLsl(t *testing.T) {
	cases := []IdentifyCase{
		{FetchedInstr16(0x0000), true}, // lsl r0, r0, #0
		{FetchedInstr16(0x01e7), true}, // lsl r7, r4, #7
		{FetchedInstr16(0x4080), true}, // lsl r0, r0, r0
		{FetchedInstr16(0x40bf), true}, // lsl r7, r7, r7
		{FetchedInstr16(0x09e7), false},// lsr r7, r4, #7
		{FetchedInstr16(0xffff), false},
	}

	instr_type := reflect.TypeOf(Lsl{})

	for _, test := range cases {
		instr, err := test.instr.Decode()
		verify_identify(t, test, instr_type, instr, err)
	}
}

func TestDecodeLslImm16(t *testing.T) {
	cases := []DecodeCase{
        // lsl r0, r0, #0
		{FetchedInstr16(0x0000), Lsl{Rd: 0, Rm: 0, Imm: 0, S: true, Rn: 0}},
        // lsl r0, r0, #1
		{FetchedInstr16(0x0040), Lsl{Rd: 0, Rm: 0, Imm: 1, S: true, Rn: 0}},
        // lsl r7, r4, #7
		{FetchedInstr16(0x01e7), Lsl{Rd: 7, Rm: 4, Imm: 7, S: true, Rn: 0}},
	}

	for _, test := range cases {
		actual := LslImm16(test.instr)
		verify_decode(t, test, actual)
	}
}

func TestDecodeLslReg16(t *testing.T) {
	cases := []DecodeCase{
        // lsl r0, r0, r0
        {FetchedInstr16(0x4080), Lsl{Rd: 0, Rn: 0, Rm: 0, Imm: 0, S: true}},
        // lsl r7, r7, r7
        {FetchedInstr16(0x40bf), Lsl{Rd: 7, Rn: 7, Rm: 7, Imm: 0, S: true}},
        // lsl r4, r4, r7
        {FetchedInstr16(0x40bc), Lsl{Rd: 4, Rn: 4, Rm: 7, Imm: 0, S: true}},
	}

	for _, test := range cases {
		actual := LslReg16(test.instr)
		verify_decode(t, test, actual)
	}
}

func TestIdentifyLsr(t *testing.T) {
	cases := []IdentifyCase{
		{FetchedInstr16(0x0800), true}, // lsr r0, r0, #0
		{FetchedInstr16(0x09e7), true}, // lsr r7, r4, #7
		{FetchedInstr16(0x40c0), true}, // lsr r0, r0, r0
		{FetchedInstr16(0x40d3), true}, // lsr r3, r3, r1
		{FetchedInstr16(0x01e7), false},// lsl r7, r4, #7
		{FetchedInstr16(0x0000), false},// lsl r0, r0, #0
		{FetchedInstr16(0xffff), false},
	}

	instr_type := reflect.TypeOf(Lsr{})

	for _, test := range cases {
		instr, err := test.instr.Decode()
		verify_identify(t, test, instr_type, instr, err)
	}
}

func TestDecodeLsrImm16(t *testing.T) {
	cases := []DecodeCase{
        // lsr r0, r0, #0
		{FetchedInstr16(0x0800), Lsr{Rd: 0, Rm: 0, Imm: 0, S: true, Rn: 0}},
        // lsr r0, r0, #1
		{FetchedInstr16(0x0840), Lsr{Rd: 0, Rm: 0, Imm: 1, S: true, Rn: 0}},
        // lsr r7, r4, #7
		{FetchedInstr16(0x09e7), Lsr{Rd: 7, Rm: 4, Imm: 7, S: true, Rn: 0}},
	}

	for _, test := range cases {
		actual := LsrImm16(test.instr)
		verify_decode(t, test, actual)
	}
}

func TestDecodeLsrReg16(t *testing.T) {
	cases := []DecodeCase{
        // lsr r0, r0, r0
        {FetchedInstr16(0x40c0), Lsr{Rd: 0, Rn: 0, Rm: 0, Imm: 0, S: true}},
        // lsr r7, r7, r7
        {FetchedInstr16(0x40ff), Lsr{Rd: 7, Rn: 7, Rm: 7, Imm: 0, S: true}},
        // lsr r4, r4, r7
        {FetchedInstr16(0x40fc), Lsr{Rd: 4, Rn: 4, Rm: 7, Imm: 0, S: true}},
	}

	for _, test := range cases {
		actual := LsrReg16(test.instr)
		verify_decode(t, test, actual)
	}
}
