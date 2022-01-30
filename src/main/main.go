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

// The value by which the display should be scaled.
const DISPLAY_SCALE = 20

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

	return nil
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

	fileName := getFileName()
	loadProgram(fileName, memory)

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	defer sdl.Quit()

	window, err := sdl.CreateWindow("CHIP-8 Interpreter",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		components.DISPLAY_WIDTH*DISPLAY_SCALE,
		components.DISPLAY_HEIGHT*DISPLAY_SCALE, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	defer window.Destroy()

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
