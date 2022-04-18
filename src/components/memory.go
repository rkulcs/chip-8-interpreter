package components

import (
	"errors"
	"fmt"
	"io"
	"os"
)

//=== Constants ===//

// The size of the buffer used when reading a CHIP-8 file's contents.
const READER_BUFFER_SIZE = 1

// The size of the memory unit used.
const MEM_SIZE = 4096

// The address of the first instruction of the program to execute.
const PROGRAM_START_LOCATION = 0x200

// The address from which the font data should be stored
const FONT_START_LOCATION = 0x50

//=== Struct Definitions ===//

type Memory struct {
	Words [MEM_SIZE]byte
}

//=== Memory Functions ===//

// Creates a new memory unit and loads the contents of the specified file
// into it.
func InitMemory(fileName string) Memory {
	memory := Memory{}
	memory.loadFont()
	memory.loadProgram(fileName)
	return memory
}

// Reads the value stored in memory at the specified address.
func (memory *Memory) ReadFrom(address int) (byte, error) {
	if address < 0 || address >= MEM_SIZE {
		return 0, errors.New("Invalid memory address.")
	}

	return memory.Words[address], nil
}

// Updates the contents of the register at the specified address to
// the given value.
func (memory *Memory) WriteTo(address int, data byte) error {
	if address < 0 || address >= MEM_SIZE {
		return errors.New("Invalid memory address.")
	}

	memory.Words[address] = data
	return nil
}

func (memory *Memory) loadFont() {

	currentAddress := FONT_START_LOCATION

	for digit := 0; digit < len(Font); digit++ {
		for i := 0; i < len(Font[digit]); i++ {
			memory.WriteTo(currentAddress, Font[digit][i])
			currentAddress++
		}
	}
}

// Loads the provided CHIP-8 file into the memory component.
func (memory *Memory) loadProgram(filePath string) error {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("The specified file does not exist.")
		os.Exit(1)
	}

	buffer := make([]byte, READER_BUFFER_SIZE)
	numBytesRead := -1

	currentAddress := PROGRAM_START_LOCATION

	for numBytesRead != 0 {
		numBytesRead, err = file.Read(buffer)

		if err != nil && err != io.EOF {
			fmt.Println("Could not finish reading the file's contents.")
			os.Exit(1)
		}

		memory.WriteTo(currentAddress, byte(buffer[0]))
		currentAddress++
	}

	_, err = memory.ReadFrom(PROGRAM_START_LOCATION)
	return err
}
