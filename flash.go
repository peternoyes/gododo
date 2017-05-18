package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"strconv"

	"go.bug.st/serial.v1"
)

func choosePort(reader *bufio.Reader) (string, error) {
	ports, err := serial.GetPortsList()
	if err != nil {
		return "", err
	}

	if len(ports) == 0 {
		return "", errors.New("No serial ports found")
	}

	if len(ports) == 1 {
		return ports[0], nil
	}

	fmt.Println("Please Select a Port:")
	for i, s := range ports {
		fmt.Printf("[%v]: %v", i, s)
		fmt.Println("")
	}
	a, _ := reader.ReadString('\n')
	a = strings.TrimSuffix(a, "\n")
	i, err := strconv.ParseInt(a, 10, 32)
	if err != nil {
		return "", err
	}

	if i < 0 || i >= int64(len(ports)) {
		return "", errors.New("Invalid Selection")
	}

	return ports[i], nil
}

func checkVersion(port string) {
	mode := &serial.Mode{BaudRate: 9600}
	s, err := serial.Open(port, mode)

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
	mode := &serial.Mode{BaudRate: 9600}
	s, err := serial.Open(port, mode)

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
	mode := &serial.Mode{BaudRate: 9600}
	s, err := serial.Open(port, mode)

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

func readByte(s serial.Port) (byte, error) {
	buf := make([]byte, 1)
	n, err := s.Read(buf)
	if err != nil {
		return 0, err
	}

	if n != 1 {
		return 0, errors.New("Did not read byte")
	}

	return buf[0], nil
}

func Flash() {
	reader := bufio.NewReader(os.Stdin)

	port, err := choosePort(reader)
	if err != nil {
		log.Fatal(err)
		return
	}

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
}
