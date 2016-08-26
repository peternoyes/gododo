package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/peternoyes/dodo-sim"
	"html/template"
	"log"
	"net/http"
	"time"
)

const (
	pongWait   = 20 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

var upgrader = websocket.Upgrader{}

func Server() {
	http.HandleFunc("/stream", stream)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

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

	c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	input := make(chan string)
	go func(input chan string) {
		for {
			_, p, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					fmt.Println("Unexpected Closer")
				}
				break
			}
			input <- string(p)
		}
	}(input)

	s := new(dodosim.Simulator)

	renderer := new(SocketRenderer)
	renderer.Conn = c
	s.Renderer = renderer

	s.Ticker = time.NewTicker(pingPeriod)
	s.Input = input
	s.IntervalCallback = func() bool {
		if err = c.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			return false
		}
		return true
	}
	s.Complete = func(cpu *dodosim.Cpu) {

	}
	s.CyclesPerFrame = func(cycles uint64) {

	}

	dodosim.Simulate(s)

	fmt.Println("Complete...")
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/stream")
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
    	if (keyState[66]) s += "B";
    	ws.send(s);
    }

    var processdown = function(evt) {
    	switch (evt.keyCode) {
    	case 37:
    	case 39:
    	case 38:
    	case 40:
    	case 65:
    	case 66:
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
    	case 66:	
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
