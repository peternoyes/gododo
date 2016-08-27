# gododo
CLI tool for the [Dodo](https://github.com/peternoyes/dodo) the 6502 Portable Game System

This tool provides several options:

1. -c will run a local simulator in the console, loading 'firmware' and 'fram.bin' from the local folder
2. -s will run a the simulator through a wepage accesible from localhost:3000
3. -f will flash fram.bin over the serial port to a connected Dodo System (note that the hardcoded USB device may need to be changed)

The code for the simulator is found int the [dodo-sim](https://github.com/peternoyes/dodo-sim) repository.

When running the simulator in the console, to exit, hit 'x'. The arrow keys as well as 'a' map to Dodo's gamepad. A keypress toggles the state of the key, due to limitations of running in the console.