package main

import (
	"os"
)

func main() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "-t":
			Test()
			return
		case "-f":
			Flash()
			return
		case "-c":
			Terminal()
			return
		}
	}

	Server()
}
