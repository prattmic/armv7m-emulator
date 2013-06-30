package main

import (
	"./core"
	"flag"
	"fmt"
	"io"
	"os"
)

var execute = flag.Bool("execute", false, "Execute instructions in addition to decoding")

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Printf("ARMv7-M Emulator\n")
		fmt.Printf("usage: %s binary\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	binary := flag.Arg(0)
	file, err := os.Open(binary)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	regs := new(core.Registers)
	b := make([]byte, 2, 2)
	addr := 0
	var upper *core.FetchedInstr16 = nil

	if *execute {
		fmt.Printf("Register state:\n")
		regs.Print()
		fmt.Printf("\n")
	}

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
			fmt.Printf(" %v", fetched16)
			fetched = upper.Extend(fetched16)
			upper = nil
		} else {
			fetched = fetched16
			fmt.Printf("%x:\t%v", addr, fetched)
		}

		addr += len(b)

		instr, err := fetched.Decode()
		if err == core.ErrIncompleteInstruction {
			upper = &fetched16
			continue
		} else if err != nil {
			fmt.Printf("\t%s\n", err)
			continue
		}

		fmt.Printf("\t%s\t%#v\n", instr, instr)

		if *execute {
			instr.Execute(regs)
			fmt.Printf("Register state:\n")
			regs.Print()
			fmt.Printf("\n")
		}
	}

	file.Close()
}
