package instructions

import "components"

func Decode(instr int32, components *components.Components) {
	firstNibble := instr & 0xF000

	switch firstNibble {
	case 0x0000:
		decode0Instruction(instr, components)
		break
	case 0x1000:
		components.Registers.PC = int16(instr & 0x0FFF)
		break
	case 0x2000:
		// TODO
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
	case 0xD000:
		decodeDInstruction(instr, components)
		break
	case 0xE000:
	case 0xF000:
	}
}

func decode0Instruction(instr int32, components *components.Components) {
	if instr == 0x00E0 {
		components.Display.Clear()
	} else if instr == 0x00EE {
		// TODO
	}
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
		*vx = *vx - *vy

		if *vx > *vy {
			*vf = 1
		} else {
			*vf = 0
		}

		break
	case 0x6:
		lsb := instr & 0x000F

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
		msb := instr & 0xF000

		if msb == 0x1000 {
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

	if *vx == *vy {
		registers.PC += 2
	}
}

func decodeDInstruction(instr int32, components *components.Components) {
	n := int(instr & 0x000F)

	vx := &(components.Registers.V[(instr>>8)&0x000F])
	vy := &(components.Registers.V[(instr>>4)&0x000F])

	vf := &(components.Registers.V[0xF])
	*vf = 0

	x := *vx % 64
	y := *vy % 32

	for i := 0; i < n; i++ {
		address := int(components.Registers.I) + i
		sprite, err := components.Memory.ReadFrom(address)

		if err != nil {
			panic(err)
		}

		x = *vx % 64

		for j := 0; j < 8; j++ {
			bit := (sprite >> (7 - j)) & 0x01
			var on bool

			if bit == 1 {
				on = true
			} else {
				on = false
			}

			pixelWasOn := components.Display.Draw(int32(x), int32(y), on)

			if pixelWasOn {
				*vf = 1
			}

			x = (x + 1) % 64
		}

		y = (y + 1) % 34
	}
}

func decodeEInstruction(instr int32, components *components.Components) {

}

func decodeFInstruction(instr int32, components *components.Components) {

}
