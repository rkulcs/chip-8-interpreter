package instructions

import (
	"components"
	"testing"
)

func TestDecode3Instruction(test *testing.T) {
	registers := components.Registers{}

	registers.V[0] = 0x10
	decode3Instruction(0x3010, &registers)

	if registers.PC != 0x2 {
		test.Errorf("PC is not incremented correctly.")
	}

	registers.V[1] = 0x11
	decode3Instruction(0x3111, &registers)

	if registers.PC != 0x4 {
		test.Errorf("PC is not incremented correctly.")
	}

	decode3Instruction(0x3410, &registers)

	if registers.PC != 0x4 {
		test.Errorf("PC is incremented upon mismatched values.")
	}

	decode3Instruction(0x3F00, &registers)

	if registers.PC != 0x6 {
		test.Errorf("PC is not incremented correctly.")
	}
}

func TestDecode4Instruction(test *testing.T) {
	registers := components.Registers{}

	value := 0x1
	expectedPC := 0x2

	for i := 0; i < 16; i++ {
		instruction := 0x4000 + (i << 8) + value
		value++

		decode4Instruction(int32(instruction), &registers)

		if registers.PC != int16(expectedPC) {
			test.Errorf("PC is not incremented correctly.")
		}

		expectedPC += 0x2
	}

	decode4Instruction(0x4900, &registers)

	if registers.PC == int16(expectedPC) {
		test.Errorf("PC is incremented upon matching values.")
	}
}

func TestDecode5Instruction(test *testing.T) {
	registers := components.Registers{}

	value := byte(0x1)
	expectedPC := 0x2

	for i := 1; i < 16; i += 2 {
		registers.V[i-1] = value
		registers.V[i] = value

		value++

		instr := 0x5000 + ((i - 1) << 8) + (i << 4)

		decode5Instruction(int32(instr), &registers)

		if registers.PC != int16(expectedPC) {
			test.Errorf("PC is not incremented correctly.")
		}

		expectedPC += 0x2
	}

	decode5Instruction(0x5180, &registers)

	if registers.PC == int16(expectedPC) {
		test.Errorf("PC is incremented upon mismatched values.")
	}
}

func TestDecode6Instruction(test *testing.T) {
	registers := components.Registers{}
	value := 0x1

	for i := 0; i < 16; i++ {
		instruction := 0x6000 + (i << 8) + value
		decode6Instruction(int32(instruction), &registers)

		if registers.V[i] != byte(value) {
			test.Errorf("Value is not loaded into V%v.", i)
		}

		value++
	}
}

func TestDecode7Instruction(test *testing.T) {
	registers := components.Registers{}
	value := 0x1

	for i := 0; i < 16; i++ {
		registers.V[i] = 0xF + byte(i)

		previous := registers.V[i]

		instruction := 0x7000 + (i << 8) + value

		decode7Instruction(int32(instruction), &registers)

		if registers.V[i] != previous+byte(value) {
			test.Errorf("Value is not added to V%v.", i)
		}

		value++
	}
}

func TestDecodeEInstruction(test *testing.T) {
	registers := &components.Registers{}
	expectedPC := 0x2

	registers.V[3] = 0x4
	instr := 0xE39E
	decodeEInstruction(int32(instr), KEY_4, registers)

	if registers.PC != int16(expectedPC) {
		test.Errorf("PC is not incremented correctly.")
	}

	instr = 0xE3A1
	decodeEInstruction(int32(instr), KEY_4, registers)
	decodeEInstruction(int32(instr), KEY_A, registers)
	expectedPC += 0x2

	if registers.PC != int16(expectedPC) {
		test.Errorf("PC is not incremented correctly.")
	}
}
