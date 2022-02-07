package components

type Registers struct {
	// Program counter.
	PC int16

	// Index register.
	I int16

	// General purpose 8-bit registers.
	V [16]byte
}
