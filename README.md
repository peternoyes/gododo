# gododo
6502 Simulator for Dodo

gododo is a simulator for (Dodo)[https://github.com/peternoyes/dodo] which is a 6502 based homebrew game system. The simulator is for the most part a port of http://rubbermallet.org/fake6502.c with some 65C02 opcodes. Decimal mode is also fixed. The simulator passes the Klaus set of 6502 tests found (here)[https://github.com/Klaus2m5]

The simulator is hardcoded to use the address space layout and devices used in Dodo but it could be repurposed for other systems I suppose. Currently the I/O is all terminal based and is a bit wonky. The code could use a lot of cleanup as well. Down the road I might try to set this up as a webservice and have a webpage that interacts with it.

Also, the 6522 Via is barely simulated at all. In the real system I use the timers for sound and a consistent framerate. I have not addressed sound at all in the simulator, and I manually trigger an interrupt based upon sleeping for 50ms rather than simulating the timer. The simulator is keeping track of accurate cycle counts for the operations. In the future I want to use this to provide warnings if the logic for each frame exceeds some limit.

When run, the simulator will open 'firmware' from the current directory and simulate it. To exit, hit 'x'. The arrow keys as well as 'a' map to Dodo's gamepad. A keypress toggles the state of the key, which is the wonkiness I described above.

To test the simulator, pass -t as an argument.