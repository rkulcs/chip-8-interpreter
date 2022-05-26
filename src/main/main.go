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

func handleInput() (int, bool) {
	// Stores the virtual key code of the last key pressed
	var keyCode int

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch eventType := event.(type) {
		case *sdl.QuitEvent:
			return -1, false
			break
		case *sdl.KeyboardEvent:
			keyCode = int(eventType.Keysym.Sym)
			break
		}
	}

	return keyCode, true
}

func executeInstructions(pause bool, x byte, keyCode int, components *components.Components) (byte, bool) {
	if pause {
		key, pressed := instructions.GetInputKeyValue(keyCode)

		if pressed {
			components.Registers.V[x] = key
			x = 0
			pause = false
		}
	}

	if (components.Registers.PC < 4096) && !pause {
		firstPart, err := components.Memory.ReadFrom(int(components.Registers.PC))
		secondPart, err := components.Memory.ReadFrom(int(components.Registers.PC) + 1)
		components.Registers.PC += 0x2

		if err != nil {
			panic(err)
		}

		instruction := (int32(firstPart) << 8) + int32(secondPart)
		x, pause = instructions.Decode(instruction, components, keyCode)
	}

	if !pause {
		components.DelayTimer.Decrement()
		components.SoundTimer.Decrement()
	}

	return x, pause
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

	// Keeps track of whether the execution of the program is paused
	pause := false

	// Keeps track of the index of the V register which may be used to store
	// the value of a key pressed
	var x byte

	// The keycode of the last key pressed
	var keyCode int

	for running {
		frameStartTime := time.Now()

		keyCode, running = handleInput()
		x, pause = executeInstructions(pause, x, keyCode, &components)

		elapsedTime := float32(time.Since(frameStartTime).Seconds())

		if elapsedTime < 0.005 {
			sdl.Delay(5 - uint32(elapsedTime*1000))
		}
	}
}
