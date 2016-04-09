package main

import (
	"fmt"
	"github.com/tarm/serial"
	"io/ioutil"
	"log"
	"time"
)

func Flash() {
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
