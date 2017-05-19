# gododo
CLI tool for the [Dodo](https://github.com/peternoyes/dodo) the 6502 Portable Game System

This tool provides several options:

1. -c will run a local simulator in the console, loading 'fram.bin' from the local folder.
2. -s will run the simulator through a wepage accessible at localhost:3000
3. -f will flash either fram.bin or the system firmware over the serial port to a connected Dodo System 

The code for the simulator is found int the [dodo-sim](https://github.com/peternoyes/dodo-sim) repository.

When running the simulator in the console, to exit, hit 'x'. The arrow keys as well as 'a' and 'b' map to Dodo's gamepad. For best results, increase the key repeat rate in your system preferences.

# Installation

1. Install go from golang.org
2. At a terminal run: go get -u github.com/peternoyes/gododo