package main

import (
	"bufio"
	"components"
	"fmt"
	"instructions"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

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
	// Get the file name of the CHIP-8 program to run
	fileName := getFileName()

	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	defer sdl.Quit()

	// Initialize CHIP-8 components and load file contents into memory
	components := components.InitComponents(fileName)
	defer components.Display.Destroy()

	running := true

	for running {
		// Stores the virtual key code of the last key pressed
		var keyCode int

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch eventType := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				keyCode = int(eventType.Keysym.Sym)
				break
			}
		}

		if components.Registers.PC < 4096 {
			firstPart, err := components.Memory.ReadFrom(int(components.Registers.PC))
			secondPart, err := components.Memory.ReadFrom(int(components.Registers.PC) + 1)
			components.Registers.PC += 0x2

			if err != nil {
				panic(err)
			}

			instruction := (int32(firstPart) << 8) + int32(secondPart)
			instructions.Decode(instruction, &components, keyCode)
		}

		time.Sleep(time.Second / 60)
		components.DelayTimer.Decrement()
	}
}
