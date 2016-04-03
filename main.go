package main

import (
	"fmt"
	"github.com/tarm/serial"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func test() {
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

func flash() {
	c := &serial.Config{Name: "/dev/tty.usbserial-A6040I72", Baud: 19200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadFile("fram.bin")
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1)
	n, err := s.Read(buf)

	if err != nil {
		log.Fatal(err)
	}

	if n != 1 {
		panic("Did not read a byte")
	}

	if buf[0] != byte('R') {
		panic("Did not read 'R'")
	} else {
		fmt.Println("Writing...")
	}

	for _, b := range dat {
		time.Sleep(1 * time.Millisecond)
		n, err := s.Write([]byte{b})
		if n != 1 {
			panic("Did not write a byte")
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done!")
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-t" {
		test()
		return
	}

	if len(os.Args) == 2 && os.Args[1] == "-f" {
		flash()
		return
	}

	bus := new(Bus)
	bus.New()

	ram := new(Ram)
	bus.Add(ram)

	rom := new(Rom)
	dat, err := ioutil.ReadFile("firmware")
	if err != nil {
		fmt.Println(err)
		return
	}

	ssd1305 := new(Ssd1305)
	ssd1305.New(ram)
	bus.Add(ssd1305)

	gamepad := new(Gamepad)
	gamepad.New()

	fram := new(Fram)
	fram.New()

	via := new(Via)
	via.New(gamepad, fram)
	bus.Add(via)

	acia := new(Acia)
	bus.Add(acia)

	for i, b := range dat {
		rom[i] = b
	}
	bus.Add(rom)

	cpu := new(Cpu)
	cpu.Reset(bus)

	BuildTable()

	var cycles uint64 = 0

	cmd := exec.Command("/bin/stty", "raw", "-echo")
	cmd.Stdin = os.Stdin
	cmd.Run()

	input := make(chan int)
	go func(ch chan int) {

		bytes := make([]byte, 3)
		var val int

		for {
			numRead, err := os.Stdin.Read(bytes)
			if err != nil {
				close(input)
				return
			}

			if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
				// Three-character control sequence, beginning with "ESC-[".

				if bytes[2] == 65 {
					// Up
					val = 38
				} else if bytes[2] == 66 {
					// Down
					val = 40
				} else if bytes[2] == 67 {
					// Right
					val = 39
				} else if bytes[2] == 68 {
					// Left
					val = 37
				}
			} else if numRead == 1 {
				val = int(bytes[0])
			} else {
				// Two characters read??
			}

			input <- val
		}

	}(input)

	syncer := make(chan int)
	go func(syncer chan int) {
		for {
			time.Sleep(50 * time.Millisecond) // 50 ms
			syncer <- 50
		}
	}(syncer)

	var lastOp uint8 = 0
	var waitTester int = 0

	for {
		opcode := bus.Read(cpu.PC)

		cpu.PC++
		cpu.Status |= Constant
		o := GetOperation(opcode)
		c := o.Execute(cpu, bus, opcode)
		cycles += uint64(c)

		if (lastOp == 0xA5 && opcode == 0xF0) || (lastOp == 0xF0 && opcode == 0xA5) {
			waitTester++
		} else {
			waitTester = 0
		}

		// If Lda, Beq sequences happens 5 times in a row then assume we are waiting for interrupt
		if waitTester == 10 {
			fmt.Printf("\033[0;66H")
			fmt.Println("Cycles Per Frame: ", cycles, "  ")
		}

		lastOp = opcode

		select {
		case b, ok := <-input:
			if !ok {
				break
			} else {
				if b == int('a') {
					gamepad.A = !gamepad.A
				} else if b == int('x') {
					fram.Flush()
					cmd = exec.Command("/bin/stty", "-raw", "echo")
					cmd.Stdin = os.Stdin
					cmd.Run()

					fmt.Println("")
					fmt.Println("A: ", cpu.A)
					fmt.Println("Y: ", cpu.Y)
					fmt.Println("X: ", cpu.X)

					return
				} else if b == 37 {
					gamepad.L = !gamepad.L
				} else if b == 39 {
					gamepad.R = !gamepad.R
				}
			}
		case <-syncer:
			cpu.Irq(bus)
			cycles = 0
			waitTester = 0
		default:
		}
	}
}
