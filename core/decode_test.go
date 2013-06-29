package core

import (
	"reflect"
	"testing"
)

// Identify type of instruction
// Don't check decoded values
type IdentifyCase struct {
	instr       FetchedInstr
	instr_valid bool
}

// Decode instruction internals
// Given known instruction
type DecodeCase struct {
	instr   FetchedInstr
	decoded DecodedInstr
}

func verify_identify(t *testing.T, test IdentifyCase, instr_type reflect.Type, identified DecodedInstr, err error) {
	if err != nil {
		if test.instr_valid {
			t.Errorf("instr: %v, err: %v", test.instr, err)
		} else {
            /* An error is OK, because this wasn't considered a valid instruction */
			return
		}
	}

    identified_type := reflect.TypeOf(identified)
    types_match := instr_type == identified_type

    if test.instr_valid && !types_match {
        t.Errorf("instr: %v, decoded type: %T, expected type: %v", test.instr, identified, instr_type)
    }

    if !test.instr_valid && types_match {
        t.Errorf("instr: %v, decoded type: %T, expected not type: %v", test.instr, identified, instr_type)
    }
}

func verify_decode(t *testing.T, test DecodeCase, actual DecodedInstr) {
	if actual != test.decoded {
		t.Errorf("instr: %v, expected %#v, got %#v", test.instr, test.decoded, actual)
	}
}

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
