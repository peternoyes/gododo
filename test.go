package main

import (
	"fmt"
	"io/ioutil"
)

func Test() {
	bus := new(Bus)
	bus.New()

	ram := new(Ramtest)
	bus.Add(ram)

	dat, err := ioutil.ReadFile("6502_functional_test.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, b := range dat {
		ram[i] = b
	}

	cpu := new(Cpu)
	cpu.Reset(bus)

	cpu.PC = 0x400

	BuildTable()

	for {
		before := cpu.PC
		opcode := bus.Read(cpu.PC)

		cpu.PC++
		cpu.Status |= Constant
		o := GetOperation(opcode)
		o.Execute(cpu, bus, opcode)

		if before == cpu.PC {
			if cpu.PC == 13209 {
				fmt.Println("Success!")
			} else {
				fmt.Println("Yikes! Trapped at: ", cpu.PC)
			}
			return
		}
	}
}
