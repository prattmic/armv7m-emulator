package main

import (
	"./core"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ARMv7-M Emulator\n")
		fmt.Printf("usage: %s binary\n", os.Args[0])
		os.Exit(1)
	}

	binary := os.Args[1]
	file, err := os.Open(binary)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	//regs := new(core.Registers)
	b := make([]byte, 2, 2)
	var upper *core.FetchedInstr16 = nil

	for {
		n, err := file.Read(b)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		} else if n != len(b) {
			fmt.Printf("not enough bytes read\n")
			os.Exit(1)
		}

		var fetched core.FetchedInstr

		fetched16 := core.FetchedInstr16((uint16(b[1]) << 8) | uint16(b[0]))

		if upper != nil {
			fetched = upper.Extend(fetched16)
			upper = nil
		} else {
			fetched = fetched16
		}

		instr, err := fetched.Decode()
		if err == core.IncompleteInstruction {
			upper = &fetched16
			continue
		} else if err != nil {
			fmt.Printf("%s (instr = %v)\n", err, fetched)
			continue
		}

		fmt.Printf("%#v\n", instr)
	}

	file.Close()
}
