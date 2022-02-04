package decoder

func decode(instr int32) {
	firstNibble := instr & 0x1000

	switch firstNibble {
	case 0x0000:
		decode0Instruction(instr)
		break
	case 0x1000:

		break
	case 0x2000:
	case 0x3000:
	case 0x4000:
	case 0x5000:
	case 0x6000:
	case 0x7000:
	case 0x8000:
	case 0x9000:
	case 0xA000:
	case 0xB000:
	case 0xC000:
	case 0xD000:
	case 0xE000:
	case 0xF000:
	}
}

func decode0Instruction(instr int32) {

}

func decode8Instruction(instr int32) {

}

func decodeEInstruction(instr int32) {

}

func decodeFInstruction(instr int32) {

}
