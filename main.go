package main

import (
    "./core"
    "fmt"
)

func main() {
    var instr core.DecodedInstr

    regs := new(core.Registers)

    regs.R[0] = 1

    fmt.Printf("%#v\n", regs)

    unknown_instr1 := core.FetchedInstr16(0xf140)
    unknown_instr2 := core.FetchedInstr16(0x0)

    instr, err := unknown_instr1.Decode()
    if (err != nil) {
        if (err == core.IncompleteInstruction) {
            unknown_instr := unknown_instr1.Extend(unknown_instr2)
            instr, err = unknown_instr.Decode()
            if (err != nil) {
                fmt.Printf("full instr: %s\n", err)
            }
        } else {
            fmt.Printf("%s\n", err)
        }
    }

    fmt.Printf("%T%+v\n", instr, instr)

    // lsl r0, r0, #1
    unknown_instr1 = core.FetchedInstr16(0x0040)

    instr, err = unknown_instr1.Decode()
    if (err != nil) {
        fmt.Printf("%s\n", err)
    }

    fmt.Printf("%T%+v\n", instr, instr)

    instr.Execute(regs)

    fmt.Printf("%#v\n", regs)

    // lsr r0, r0, #1
    unknown_instr1 = core.FetchedInstr16(0x0840)

    instr, err = unknown_instr1.Decode()
    if (err != nil) {
        fmt.Printf("%s\n", err)
    }

    fmt.Printf("%T%+v\n", instr, instr)

    instr.Execute(regs)

    fmt.Printf("%#v\n", regs)
}
