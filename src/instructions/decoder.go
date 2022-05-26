package instructions

import (
	"components"
	"math/rand"
	"time"
)

func Decode(instr int32, components *components.Components, keyCode int) (x byte, pause bool) {
	firstNibble := instr & 0xF000

	switch firstNibble {
	case 0x0000:
		decode0Instruction(instr, components)
		break
	case 0x1000:
		components.Registers.PC = int16(instr & 0x0FFF)
		break
	case 0x2000:
		decode2Instruction(instr, components)
		break
	case 0x3000:
		decode3Instruction(instr, components.Registers)
		break
	case 0x4000:
		decode4Instruction(instr, components.Registers)
		break
	case 0x5000:
		decode5Instruction(instr, components.Registers)
		break
	case 0x6000:
		decode6Instruction(instr, components.Registers)
		break
	case 0x7000:
		decode7Instruction(instr, components.Registers)
		break
	case 0x8000:
		decode8Instruction(instr, components.Registers)
		break
	case 0x9000:
		decode9Instruction(instr, components.Registers)
		break
	case 0xA000:
		components.Registers.I = int16(instr & 0x0FFF)
		break
	case 0xB000:
		components.Registers.PC = int16(instr&0x0FFF) +
			int16(components.Registers.V[0x0])
		break
	case 0xC000:
		decodeCInstruction(instr, components.Registers)
		break
	case 0xD000:
		decodeDInstruction(instr, components)
		break
	case 0xE000:
		decodeEInstruction(instr, keyCode, components.Registers)
		break
	case 0xF000:
		x, pause := decodeFInstruction(instr, components)
		return x, pause
	}

	return 0, false
}

func decode0Instruction(instr int32, components *components.Components) {
	if instr == 0x00E0 {
		components.Display.Clear()
	} else if instr == 0x00EE {
		var err error
		components.Registers.PC, err = components.Stack.Pop()

		if err != nil {
			panic(err)
		}
	}
}

func decode2Instruction(instr int32, components *components.Components) {
	components.Stack.Push(components.Registers.PC)
	components.Registers.PC = int16(instr & 0x0FFF)
}

func decode3Instruction(instr int32, registers *components.Registers) {
	registerIndex := (instr >> 8) & 0x000F
	valueToMatch := byte(instr & 0x00FF)

	if registers.V[registerIndex] == valueToMatch {
		registers.PC += 2
	}
}

func decode4Instruction(instr int32, registers *components.Registers) {
	registerIndex := (instr >> 8) & 0x000F
	valueToMatch := byte(instr & 0x00FF)

	if registers.V[registerIndex] != valueToMatch {
		registers.PC += 2
	}
}

func decode5Instruction(instr int32, registers *components.Registers) {
	x := (instr >> 8) & 0x000F
	y := (instr >> 4) & 0x000F

	if registers.V[x] == registers.V[y] {
		registers.PC += 2
	}
}

func decode6Instruction(instr int32, registers *components.Registers) {
	registerIndex := (instr >> 8) & 0x000F
	value := byte(instr & 0x00FF)

	registers.V[registerIndex] = value
}

func decode7Instruction(instr int32, registers *components.Registers) {
	registerIndex := (instr >> 8) & 0x000F
	value := byte(instr & 0x00FF)

	registers.V[registerIndex] += value
}

func decode8Instruction(instr int32, registers *components.Registers) {
	x := (instr >> 8) & 0x000F
	y := (instr >> 4) & 0x000F

	vx := &(registers.V[x])
	vy := &(registers.V[y])
	vf := &(registers.V[0xF])
	op := instr & 0x000F

	switch op {
	case 0x0:
		*vx = *vy
		break
	case 0x1:
		*vx = *vx | *vy
		break
	case 0x2:
		*vx = *vx & *vy
		break
	case 0x3:
		*vx = *vx ^ *vy
		break
	case 0x4:
		comp := int16(*vx) + int16(*vy)
		*vx = *vx + *vy

		if int16(*vx) != comp {
			*vf = 1
		} else {
			*vf = 0
		}

		break
	case 0x5:
		if *vx > *vy {
			*vf = 1
		} else {
			*vf = 0
		}

		*vx = *vx - *vy

		break
	case 0x6:
		lsb := *vx & 0x0001

		if lsb == 1 {
			*vf = 1
		} else {
			*vf = 0
		}

		*vx /= 0x2
		break
	case 0x7:
		*vx = *vy - *vx

		if *vy > *vx {
			*vf = 1
		} else {
			*vf = 0
		}

		break
	case 0xE:
		msb := *vx >> 7

		if msb == 1 {
			*vf = 1
		} else {
			*vf = 0
		}

		*vx *= 0x2
		break
	}
}

func decode9Instruction(instr int32, registers *components.Registers) {
	x := (instr >> 8) & 0x000F
	y := (instr >> 4) & 0x000F

	vx := &(registers.V[x])
	vy := &(registers.V[y])

	if *vx != *vy {
		registers.PC += 2
	}
}

func decodeCInstruction(instr int32, registers *components.Registers) {
	x := (instr >> 8) & 0x000F
	vx := &(registers.V[x])

	rand.Seed(time.Now().UnixNano())

	var randomByte byte = byte(rand.Intn(256))
	*vx = randomByte & byte(instr&0x00FF)
}

func decodeDInstruction(instr int32, components *components.Components) {
	// Get the number of bytes to read sprite data from
	n := int(instr & 0x000F)

	vx := &(components.Registers.V[(instr>>8)&0x000F])
	vy := &(components.Registers.V[(instr>>4)&0x000F])

	vf := &(components.Registers.V[0xF])
	*vf = 0

	x := *vx % 64
	y := *vy % 32

	for i := 0; i < n; i++ {
		// Get the next byte containing sprite data
		address := int(components.Registers.I) + i
		sprite, err := components.Memory.ReadFrom(address)

		if err != nil {
			panic(err)
		}

		// Reset the value of x
		x = *vx % 64

		for j := 0; j < 8; j++ {
			// Read the next bit from the sprite
			bit := (sprite >> (7 - j)) & 0x01
			var on bool

			if bit == 1 {
				on = true
			} else {
				on = false
			}

			// Draw the sprite onto the screen
			pixelWasOn := components.Display.Draw(int32(x), int32(y), on)

			// Set VF to 1 if there was a white pixel at the current coordinates
			if pixelWasOn && on {
				*vf = 1
			}

			if x == 63 {
				break
			} else {
				x++
			}
		}

		if y == 31 {
			break
		} else {
			y++
		}
	}
}

func decodeEInstruction(instr int32, keyCode int, registers *components.Registers) {
	vx := &(registers.V[(instr>>8)&0x000F])
	op := instr & 0x00FF

	if keyCode == 0 {
		return
	}

	key, pressed := GetInputKeyValue(keyCode)

	if pressed && (op == 0x9E) && (*vx == key) {
		registers.PC += 0x2
	}

	if (op == 0xA1) && (*vx != key) {
		registers.PC += 0x2
	}
}

func decodeFInstruction(instr int32, components *components.Components) (x byte, pause bool) {
	x = byte((instr >> 8) & 0x000F)
	vx := &(components.Registers.V[x])
	op := instr & 0x00FF

	switch op {
	case 0x07:
		*vx = components.DelayTimer.Value
		break
	case 0x0A:
		// Signal that the program needs to be paused
		pause = true
		return x, pause
	case 0x15:
		components.DelayTimer.Value = *vx
		break
	case 0x18:
		components.SoundTimer.Value = *vx
		break
	case 0x1E:
		components.Registers.I += int16(*vx)
		break
	case 0x29:
		components.Registers.I = int16(components.Display.GetFontLocation(*vx))
		break
	case 0x33:
		storeBCD(vx, components)
		break
	case 0x55:
		storeV(components)
		break
	case 0x65:
		loadV(components)
		break
	}

	return 0, false
}

func storeBCD(vx *byte, components *components.Components) {
	value := *vx
	currentAddress := components.Registers.I

	var div byte = 100

	for div > 0 {
		digit := value / div
		components.Memory.WriteTo(int(currentAddress), digit)
		currentAddress++
		value -= (digit * div)
		div /= 10
	}
}

func storeV(components *components.Components) {
	currentAddress := components.Registers.I

	for i := 0; i < 16; i++ {
		components.Memory.WriteTo(int(currentAddress), components.Registers.V[i])
		currentAddress++
	}
}

func loadV(components *components.Components) {
	currentAddress := components.Registers.I

	for i := 0; i < 16; i++ {
		value, err := components.Memory.ReadFrom((int(currentAddress)))

		if err != nil {
			panic("Unable to read memory contents")
		}

		components.Registers.V[i] = value
		currentAddress++
	}
}
