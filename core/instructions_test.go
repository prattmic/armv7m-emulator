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

// Execute given instruction on
// set of registers, with expected
// result
type ExecuteCase struct {
	instr    DecodedInstr
	regs     Registers
	expected Registers
}

func test_identify(t *testing.T, cases []IdentifyCase, instr_type reflect.Type) {
	for _, test := range cases {
		instr, err := test.instr.Decode()
		verify_identify(t, test, instr_type, instr, err)
	}
}

func test_decode(t *testing.T, cases []DecodeCase, decode DecodeFunc) {
	for _, test := range cases {
		actual := decode(test.instr)
		if actual != test.decoded {
			t.Errorf("instr: %#v", test.instr)
			t.Errorf("decoded: %#v", actual)
			t.Errorf("expected: %#v", test.decoded)
		}
	}
}

func test_execute(t *testing.T, cases []ExecuteCase) {
	for _, test := range cases {
		original := test.regs
		test.instr.Execute(&test.regs)

		if test.regs != test.expected {
			t.Errorf("instr: %#v", test.instr)
			t.Errorf("Before:\n%s", original.Pretty())
			t.Errorf("After:\n%s", test.regs.Pretty())
			t.Errorf("Expected:\n%s", test.expected.Pretty())
		}
	}
}

func verify_identify(t *testing.T, test IdentifyCase, instr_type reflect.Type, identified DecodedInstr, err error) {
	if err != nil {
		if test.instr_valid {
			t.Errorf("instr: %#v", test.instr)
			t.Errorf("err: %v", err)
			return
		} else {
			/* An error is OK, because this wasn't considered a valid instruction */
			return
		}
	}

	identified_type := reflect.TypeOf(identified)
	types_match := instr_type == identified_type

	if test.instr_valid && !types_match {
		t.Errorf("instr: %#v", test.instr)
		t.Errorf("decoded type: %T", identified)
		t.Errorf("expected type: %v", instr_type)
	}

	if !test.instr_valid && types_match {
		t.Errorf("instr: %#v", test.instr)
		t.Errorf("decoded type: %T", identified)
		t.Errorf("expected NOT type: %v", instr_type)
	}
}
