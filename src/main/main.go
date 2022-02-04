package main

import (
	"bufio"
	"components"
	"fmt"
	"io"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// The size of the buffer used when reading a CHIP-8 file's contents.
const READER_BUFFER_SIZE = 1

// Loads the provided CHIP-8 file into the memory component.
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

	_, err = memory.ReadFrom(components.PROGRAM_START_LOCATION)
	return err
}

// Prompts to the user to enter the name of a CHIP-8 file.
// Returns the name of the file.
func getFileName() string {
	inputReader := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a CHIP-8 file name: ")
	inputReader.Scan()
	fileName := inputReader.Text()

	return fileName
}

func main() {
	memory := &components.Memory{}

	// General purpose 8-bit registers
	// v := [0xF]byte{}

	// Program counter
	// var pc int16

	fileName := getFileName()
	loadProgram(fileName, memory)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	defer sdl.Quit()

	display := components.InitDisplay()
	defer display.Destroy()

	running := true

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}
}
