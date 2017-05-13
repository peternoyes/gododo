package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/tarm/serial"
)

func checkVersion(port string) {
	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)

	defer s.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sent command to read version")
	n, err := s.Write([]byte{byte('V')})
	if n != 1 {
		panic("Did not write a byte")
	}
	if err != nil {
		log.Fatal(err)
	}

	version := make([]byte, 0)
	for {
		r, err := readByte(s)

		fmt.Println("Just Read: ", r, " ", string([]byte{r}))

		if err != nil {
			log.Fatal(err)
			return
		}

		if r == 0 {
			break
		}

		version = append(version, r)
	}

	fmt.Println("Read Version: ", string(version))
}

func flashGame(port string) {
	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)

	defer s.Close()

	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadFile("fram.bin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sent command to flash game")
	n, err := s.Write([]byte{byte('G')})
	if n != 1 {
		panic("Did not write a byte")
	}
	if err != nil {
		log.Fatal(err)
	}

	r, err := readByte(s)
	if err != nil {
		log.Fatal(err)
		return
	}

	if r != byte('A') {
		panic("Did not read 'A'")
		return
	}

	for _, b := range dat {
		n, err = s.Write([]byte{b})
		if n != 1 {
			panic("Did not write a byte")
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	r, err = readByte(s)
	if err != nil {
		log.Fatal(err)
		return
	}

	if r != byte('A') {
		panic("Did not read 'A'")
		return
	}

	fmt.Println("Success!")
}

func flashSystem(port string) {
	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)

	defer s.Close()

	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadFile("firmware")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sent command to flash system")
	n, err := s.Write([]byte{byte('S')})
	if n != 1 {
		panic("Did not write a byte")
	}
	if err != nil {
		log.Fatal(err)
	}

	for i, b := range dat {
		if i%64 == 0 {
			r, err := readByte(s)
			if err != nil {
				log.Fatal(err)
				return
			}

			if r != byte('A') {
				fmt.Println("The byte is ", r)
				panic("Did not read 'A'")
			}
		}

		n, err = s.Write([]byte{b})
		if n != 1 {
			panic("Did not write a byte")
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	r, err := readByte(s)
	if err != nil {
		log.Fatal(err)
		return
	}

	if r != byte('A') {
		panic("Did not read 'A'")
		return
	}

	fmt.Println("Success!")
}

func readByte(s *serial.Port) (byte, error) {
	buf := make([]byte, 1)
	n, err := s.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != 1 {
		return 0, errors.New("Did not read single byte")
	}

	return buf[0], nil
}

func Flash() {
	port := "/dev/tty.usbserial-A6040I72"

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Flash [S]ystem or [G]ame, check [V]ersion or [Q]uit?")
		a, _ := reader.ReadString('\n')
		a = strings.TrimSuffix(a, "\n")
		switch a {
		case "s", "S":
			fmt.Println("Flashing system...")
			flashSystem(port)
			return
		case "g", "G":
			fmt.Println("Flashing game...")
			flashGame(port)
			return
		case "v", "V":
			fmt.Println("Checking version...")
			checkVersion(port)
			return
		case "q", "Q":
			fmt.Println("Quiting...")
			return
		default:
			fmt.Println("Invalid response")
			break
		}
	}

	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)

	defer s.Close()

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
		fmt.Println(rune(buf[0]))
		panic("Did not read 'R'")
	} else {
		fmt.Println("Writing...")
	}

	fmt.Println("Wrote a G")
	n, err = s.Write([]byte{byte('G')})

	if n != 1 {
		panic("Did not write a byte")
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range dat {
		n, err = s.Write([]byte{b})
		if n != 1 {
			panic("Did not write a byte")
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	s.Read(buf)

	fmt.Println("Done!")
}
