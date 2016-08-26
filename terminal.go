package main

import (
	"fmt"
	"github.com/peternoyes/dodo-sim"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Terminal() {
	cmd := exec.Command("/bin/stty", "raw", "-echo")
	cmd.Stdin = os.Stdin
	cmd.Run()

	input := make(chan string)
	go func(ch chan string) {

		bytes := make([]byte, 3)
		var state string = ""

		toggleRune := func(r rune) {
			i := strings.IndexRune(state, r)
			if i < 0 {
				state += string(r)
			} else {
				state = strings.Replace(state, string(r), "", 1)
			}
		}

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
					toggleRune('U')
				} else if bytes[2] == 66 {
					// Down
					toggleRune('D')
				} else if bytes[2] == 67 {
					// Right
					toggleRune('R')
				} else if bytes[2] == 68 {
					// Left
					toggleRune('L')
				}
			} else if numRead == 1 {
				switch rune(bytes[0]) {
				case 'a':
					fallthrough
				case 'A':
					toggleRune('A')
					break
				case 'x':
					fallthrough
				case 'X':
					toggleRune('X')
					break
				}
			} else {
				// Two characters read??
			}

			input <- state
		}

	}(input)

	s := new(dodosim.Simulator)
	s.Renderer = new(ConsoleRenderer)

	s.Ticker = time.NewTicker(1 * time.Second) // Not using
	s.Input = input
	s.IntervalCallback = func() bool { return true }
	s.Complete = func(cpu *dodosim.Cpu) {
		cmd := exec.Command("/bin/stty", "-raw", "echo")
		cmd.Stdin = os.Stdin
		cmd.Run()

		fmt.Println("")
		fmt.Println("A: ", cpu.A)
		fmt.Println("Y: ", cpu.Y)
		fmt.Println("X: ", cpu.X)
	}
	s.CyclesPerFrame = func(cycles uint64) {
		fmt.Printf("\033[0;66H")
		fmt.Println("Cycles Per Frame: ", cycles, "  ")
	}

	dodosim.Simulate(s)
}

type ConsoleRenderer struct {
}

func (r *ConsoleRenderer) Render(data [1024]byte) {
	fmt.Printf("\033[0;0H")
	var x, y int
	for y = 0; y < 64; y += 2 {
		for x = 0; x < 128; x += 2 {
			p1 := (data[x+(y/8)*128] >> uint(y%8)) & 0x1
			p2 := (data[x+1+(y/8)*128] >> uint(y%8)) & 0x1
			p3 := (data[x+((y+1)/8)*128] >> uint((y+1)%8)) & 0x1
			p4 := (data[x+1+((y+1)/8)*128] >> uint((y+1)%8)) & 0x1
			if p1 == 0x0 && p2 == 0x0 && p3 == 0x0 && p4 == 0x0 {
				fmt.Print(" ")
			} else if p1 == 0x0 && p2 == 0x0 && p3 == 0x0 && p4 == 0x1 {
				fmt.Print("\u2597")
			} else if p1 == 0x0 && p2 == 0x0 && p3 == 0x1 && p4 == 0x0 {
				fmt.Print("\u2596")
			} else if p1 == 0x0 && p2 == 0x0 && p3 == 0x1 && p4 == 0x1 {
				fmt.Print("\u2584")
			} else if p1 == 0x0 && p2 == 0x1 && p3 == 0x0 && p4 == 0x0 {
				fmt.Print("\u259D")
			} else if p1 == 0x0 && p2 == 0x1 && p3 == 0x0 && p4 == 0x1 {
				fmt.Print("\u2590")
			} else if p1 == 0x0 && p2 == 0x1 && p3 == 0x1 && p4 == 0x0 {
				fmt.Print("\u259E")
			} else if p1 == 0x0 && p2 == 0x1 && p3 == 0x1 && p4 == 0x1 {
				fmt.Print("\u259F")
			} else if p1 == 0x1 && p2 == 0x0 && p3 == 0x0 && p4 == 0x0 {
				fmt.Print("\u2598")
			} else if p1 == 0x1 && p2 == 0x0 && p3 == 0x0 && p4 == 0x1 {
				fmt.Print("\u259A")
			} else if p1 == 0x1 && p2 == 0x0 && p3 == 0x1 && p4 == 0x0 {
				fmt.Print("\u258C")
			} else if p1 == 0x1 && p2 == 0x0 && p3 == 0x1 && p4 == 0x1 {
				fmt.Print("\u2599")
			} else if p1 == 0x1 && p2 == 0x1 && p3 == 0x0 && p4 == 0x0 {
				fmt.Print("\u2580")
			} else if p1 == 0x1 && p2 == 0x1 && p3 == 0x0 && p4 == 0x1 {
				fmt.Print("\u259C")
			} else if p1 == 0x1 && p2 == 0x1 && p3 == 0x1 && p4 == 0x0 {
				fmt.Print("\u259B")
			} else if p1 == 0x1 && p2 == 0x1 && p3 == 0x1 && p4 == 0x1 {
				fmt.Print("\u2588")
			}
		}
		fmt.Print("\r\n")
	}
}
