package main

import (
	"fmt"
	"io/ioutil"

	"time"

	"strconv"

	termbox "github.com/nsf/termbox-go"
	"github.com/peternoyes/dodo-sim"
)

func Terminal() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	keys := new(Keys)
	keys.New()

	firmware, err := Asset("data/firmware")
	if err != nil {
		fmt.Println(err)
		return
	}

	game, err := ioutil.ReadFile("fram.bin")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := new(dodosim.SimulatorSync)

	s.CyclesPerFrame = func(cycles uint64) {
		drawString(66, 0, "Cycles/Frame: "+strconv.Itoa(int(cycles))+"  ")
	}

	s.Renderer = new(ConsoleRenderer)
	s.SimulateSyncInit(firmware, game)

	c := time.NewTicker(50 * time.Millisecond).C

	for !keys.Done {

		state := ""
		if keys.IsPressed(A) {
			state += "A"
		}
		if keys.IsPressed(B) {
			state += "B"
		}
		if keys.IsPressed(Up) {
			state += "U"
		}
		if keys.IsPressed(Left) {
			state += "L"
		}
		if keys.IsPressed(Right) {
			state += "R"
		}
		if keys.IsPressed(Down) {
			state += "D"
		}

		s.PumpClock(state)

		<-c
	}
}

type ConsoleRenderer struct {
}

func drawString(x, y int, text string) {
	for _, s := range text {
		termbox.SetCell(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}

	termbox.Flush()
}

func (r *ConsoleRenderer) Render(data [1024]byte) {
	var x, y int
	for y = 0; y < 64; y += 2 {
		for x = 0; x < 128; x += 2 {
			p1 := (data[x+(y/8)*128] >> uint(y%8)) & 0x1
			p2 := (data[x+1+(y/8)*128] >> uint(y%8)) & 0x1
			p3 := (data[x+((y+1)/8)*128] >> uint((y+1)%8)) & 0x1
			p4 := (data[x+1+((y+1)/8)*128] >> uint((y+1)%8)) & 0x1
			if p1 == 0x0 && p2 == 0x0 && p3 == 0x0 && p4 == 0x0 {
				termbox.SetCell(x/2, y/2, ' ', termbox.ColorBlack, termbox.ColorBlack)
			} else if p1 == 0x0 && p2 == 0x0 && p3 == 0x0 && p4 == 0x1 {
				termbox.SetCell(x/2, y/2, '\u2597', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x0 && p2 == 0x0 && p3 == 0x1 && p4 == 0x0 {
				termbox.SetCell(x/2, y/2, '\u2596', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x0 && p2 == 0x0 && p3 == 0x1 && p4 == 0x1 {
				termbox.SetCell(x/2, y/2, '\u2584', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x0 && p2 == 0x1 && p3 == 0x0 && p4 == 0x0 {
				termbox.SetCell(x/2, y/2, '\u259D', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x0 && p2 == 0x1 && p3 == 0x0 && p4 == 0x1 {
				termbox.SetCell(x/2, y/2, '\u2590', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x0 && p2 == 0x1 && p3 == 0x1 && p4 == 0x0 {
				termbox.SetCell(x/2, y/2, '\u259E', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x0 && p2 == 0x1 && p3 == 0x1 && p4 == 0x1 {
				termbox.SetCell(x/2, y/2, '\u259F', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x1 && p2 == 0x0 && p3 == 0x0 && p4 == 0x0 {
				termbox.SetCell(x/2, y/2, '\u2598', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x1 && p2 == 0x0 && p3 == 0x0 && p4 == 0x1 {
				termbox.SetCell(x/2, y/2, '\u259A', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x1 && p2 == 0x0 && p3 == 0x1 && p4 == 0x0 {
				termbox.SetCell(x/2, y/2, '\u258C', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x1 && p2 == 0x0 && p3 == 0x1 && p4 == 0x1 {
				termbox.SetCell(x/2, y/2, '\u2599', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x1 && p2 == 0x1 && p3 == 0x0 && p4 == 0x0 {
				termbox.SetCell(x/2, y/2, '\u2580', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x1 && p2 == 0x1 && p3 == 0x0 && p4 == 0x1 {
				termbox.SetCell(x/2, y/2, '\u259C', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x1 && p2 == 0x1 && p3 == 0x1 && p4 == 0x0 {
				termbox.SetCell(x/2, y/2, '\u259B', termbox.ColorGreen, termbox.ColorBlack)
			} else if p1 == 0x1 && p2 == 0x1 && p3 == 0x1 && p4 == 0x1 {
				termbox.SetCell(x/2, y/2, '\u2588', termbox.ColorGreen, termbox.ColorBlack)
			}
		}
		termbox.Flush()
	}
}
