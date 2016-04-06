package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/tarm/serial"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
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

var upgrader = websocket.Upgrader{}

type SocketRenderer struct {
	Conn *websocket.Conn
}

func (s *SocketRenderer) Render(data [1024]byte) {
	s.Conn.WriteMessage(websocket.BinaryMessage, data[:])
}

func stream(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade: ", err)
		return
	}
	defer c.Close()

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

	renderer := new(SocketRenderer)
	renderer.Conn = c

	ssd1305 := new(Ssd1305)
	ssd1305.New(ram, renderer)

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

	syncer := make(chan int)
	go func(syncer chan int) {
		for {
			time.Sleep(50 * time.Millisecond) // 50 ms
			syncer <- 50
		}
	}(syncer)

	input := make(chan string)
	go func(input chan string) {
		for {
			_, p, err := c.ReadMessage()
			if err != nil {
				return
			}
			input <- string(p)
		}
	}(input)

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
		//if waitTester == 10 {
		//	fmt.Printf("\033[0;66H")
		//	fmt.Println("Cycles Per Frame: ", cycles, "  ")
		//}

		lastOp = opcode

		select {
		case <-syncer:
			cpu.Irq(bus)
			cycles = 0
			waitTester = 0

		case s, ok := <-input:
			if !ok {
				break
			} else {
				gamepad.A = strings.Contains(s, "A")
				gamepad.L = strings.Contains(s, "L")
				gamepad.R = strings.Contains(s, "R")
			}
		default:
		}
	}

	fmt.Println("Socket Done")
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/stream")
}

func server() {
	http.HandleFunc("/stream", stream)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
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

	if len(os.Args) == 2 && os.Args[1] == "-s" {
		server()
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

	consoleRenderer := new(ConsoleRenderer)

	ssd1305 := new(Ssd1305)
	ssd1305.New(ram, consoleRenderer)

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

var homeTemplate = template.Must(template.New("").Parse(`<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>
window.addEventListener("load", function(evt) {
	var ws;
	var output = document.getElementById("output");
	var c = document.getElementById("myCanvas");
	var ctx = c.getContext("2d");
	var on = ctx.createImageData(1, 1);
	var off = ctx.createImageData(1, 1);
	var d = on.data;
	d[0] = 255;
	d[1] = 255;
	d[2] = 255;
	d[3] = 255;
	d = off.data;
	d[0] = 0;
	d[1] = 0;
	d[2] = 0;
	d[3] = 255;

	var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    var keyState = {};

    var sendKeyState = function() {
    	var s = "";
    	if (keyState[37]) s += "L";
    	if (keyState[38]) s += "U";
    	if (keyState[39]) s += "R";
    	if (keyState[40]) s += "D";
    	if (keyState[65]) s += "A";
    	ws.send(s);
    }

    var processdown = function(evt) {
    	switch (evt.keyCode) {
    	case 37:
    	case 39:
    	case 38:
    	case 40:
    	case 65:
    		if (!keyState[evt.keyCode]) {
    			keyState[evt.keyCode] = true;
    			sendKeyState();
    		}
    		break;
    	}
    }

    var processup = function(evt) {
    	switch (evt.keyCode) {
    	case 37:
    	case 39:
    	case 38:
    	case 40:    
    	case 65:		
			keyState[evt.keyCode] = false;
			sendKeyState();
    		break;
    	}
    }


    window.addEventListener("keydown", processdown);
    window.addEventListener("keyup", processup);

	document.getElementById("go").onclick = function(evt) {
		ws = new WebSocket("{{.}}");
		ws.binaryType = 'arraybuffer';
		ws.onmessage = function(evt) {
			var dv = new DataView(evt.data);
			var x, y = 0;
			for (y = 0; y < 64; ++y) {
				for (x = 0; x < 128; ++x) {
					b = (dv.getUint8(x+(Math.floor(y/8)*128)) >> (y%8)) & 1;
					if (b == 1) {
						ctx.putImageData(on, x, y);
					} else {
						ctx.putImageData(off, x, y);
					}
				}
			}			
		}
		return false;
	};
});
</script>
</head>
<body>
<form>
<button id="go">Go</button>
</form>
<canvas id="myCanvas" width="128" height="64"></canvas>
<div id="output"></div>
</body>
</html>
`))
