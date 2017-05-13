package main

import termbox "github.com/nsf/termbox-go"
import "time"

type Key int

const (
	Up Key = iota
	Down
	Left
	Right
	A
	B
)

type keyState struct {
	down        bool
	tickCount   int
	repeatCount int
	stamp       time.Time
}

type Keys struct {
	state  map[Key]*keyState
	ticker <-chan time.Time
	Done   bool
}

func (k *Keys) New() {
	k.state = make(map[Key]*keyState)
	k.ticker = time.NewTicker(15 * time.Millisecond).C
	go k.watch()
	go k.wait()
}

func (k *Keys) watch() {
	for {
		select {
		case <-k.ticker:
			for _, s := range k.state {
				if s.down {
					s.tickCount++

					if s.repeatCount > 0 {
						limit := 50
						if s.repeatCount == 1 {
							limit = 250
						}

						if time.Now().Sub(s.stamp) > time.Duration(limit)*time.Millisecond {
							s.down = false
							s.repeatCount = 0
							s.tickCount = 0
						}
					}
				}
			}
		}
	}
}

func (k *Keys) wait() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				k.add(Left)
			case termbox.KeyArrowRight:
				k.add(Right)
			case termbox.KeyArrowUp:
				k.add(Up)
			case termbox.KeyArrowDown:
				k.add(Down)
			//case termbox.KeyEsc:	// This unexpectadly happens
			//	k.Done = true
			default:
				if ev.Ch == 'a' {
					k.add(A)
				} else if ev.Ch == 'b' {
					k.add(B)
				} else if ev.Ch == 'x' {
					k.Done = true
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func (k *Keys) add(key Key) {
	s, ok := k.state[key]
	if !ok {
		s = &keyState{}
		k.state[key] = s
	}

	s.down = true
	s.repeatCount++
	s.stamp = time.Now()
}

func (k *Keys) IsPressed(key Key) bool {
	s, ok := k.state[key]
	if ok {
		return s.down
	}
	return false
}
