package main

type Via struct {
	L bool
	R bool
	A bool
}

func (v *Via) New() {
	v.L = false
	v.R = false
	v.A = false
}

func (v *Via) Start() uint16 {
	return 0x6000
}

func (v *Via) Length() uint32 {
	return 0x10
}

func (v *Via) Read(addr uint16) uint8 {
	val := uint8(0xFF)
	if addr == 0x6001 {
		if v.L {
			val &^= uint8(4)
		}

		if v.R {
			val &^= uint8(8)
		}

		if v.A {
			val &^= uint8(16)
		}
	}

	return val
}

func (v *Via) Write(addr uint16, val uint8) {

}
