package components

import "errors"

const MEM_SIZE = 4096
const PROGRAM_START_LOCATION = 0x200

type Memory struct {
	Words [MEM_SIZE]byte
}

func (memory *Memory) ReadFrom(address int) (byte, error) {
	if address < 0 || address >= MEM_SIZE {
		return 0, errors.New("Invalid memory address.")
	}

	return memory.Words[address], nil
}

func (memory *Memory) WriteTo(address int, data byte) error {
	if address < 0 || address >= MEM_SIZE {
		return errors.New("Invalid memory address.")
	}

	memory.Words[address] = data
	return nil
}
