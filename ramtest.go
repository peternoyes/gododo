package main

type Ramtest [0x10000]uint8

func (ram *Ramtest) Start() uint16 {
	return 0x0
}

func (ram *Ramtest) Length() uint32 {
	return 0x10000
}

func (ram *Ramtest) Read(addr uint16) uint8 {
	return ram[addr]
}

func (ram *Ramtest) Write(addr uint16, val uint8) {
	ram[addr] = val
}
