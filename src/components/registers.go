package components

type Registers struct {
	// Program counter.
	pc int16

	// Index register.
	i int16

	// General purpose 8-bit registers.
	v [0xF]byte
}
