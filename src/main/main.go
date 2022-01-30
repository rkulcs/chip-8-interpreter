package main

import (
	"bufio"
	"components"
	"fmt"
	"io"
	"os"
)

const READER_BUFFER_SIZE = 1

func loadProgram(filePath string, memory *components.Memory) error {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("The specified file does not exist.")
		os.Exit(1)
	}

	buffer := make([]byte, READER_BUFFER_SIZE)
	numBytesRead := -1

	currentAddress := components.PROGRAM_START_LOCATION

	for numBytesRead != 0 {
		numBytesRead, err = file.Read(buffer)

		if err != nil && err != io.EOF {
			fmt.Println("Could not finish reading the file's contents.")
			os.Exit(1)
		}

		memory.WriteTo(currentAddress, byte(buffer[0]))
		currentAddress++
	}

	return nil
}

func main() {
	inputReader := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a CHIP-8 file name: ")
	inputReader.Scan()
	fileName := inputReader.Text()

	memory := &components.Memory{}

	loadProgram(fileName, memory)
}
