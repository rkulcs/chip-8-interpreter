package main

import (
	"bufio"
	"components"
	"fmt"
	"instructions"
	"os"
	"time"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

const TARGET_FPS = 120
const SFX_PATH = "../../assets/sfx.wav"

// Prompts to the user to enter the name of a CHIP-8 file.
// Returns the name of the file.
func getFileName() string {
	inputReader := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a CHIP-8 file path: ")
	inputReader.Scan()
	fileName := inputReader.Text()

	return fileName
}

func handleInput(components *components.Components) bool {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch eventType := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			keyCode := int(eventType.Keysym.Sym)

			if eventType.GetType() == sdl.KEYDOWN {
				components.InputMap.SetInputKeyState(keyCode, true)
			} else if eventType.GetType() == sdl.KEYUP {
				components.InputMap.SetInputKeyState(keyCode, false)
			}
		}
	}

	return true
}

func executeInstructions(components *components.Components) {

	if components.Registers.PC < 4096 {
		firstPart, err := components.Memory.ReadFrom(int(components.Registers.PC))
		secondPart, err := components.Memory.ReadFrom(int(components.Registers.PC) + 1)
		components.Registers.PC += 0x2

		if err != nil {
			panic(err)
		}

		instruction := (int32(firstPart) << 8) + int32(secondPart)

		instructions.Decode(instruction, components)
	}
}

func main() {
	// Get the file name of the CHIP-8 program to run
	fileName := getFileName()

	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	// Initialize audio
	if err := mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT,
		mix.DEFAULT_CHANNELS, mix.DEFAULT_CHUNKSIZE); err != nil {
		panic(err)
	}

	defer mix.CloseAudio()

	// Load sound effect
	sound, err := mix.LoadWAV(SFX_PATH)

	if err != nil {
		panic(err)
	}

	defer sound.Free()

	defer sdl.Quit()

	// Initialize CHIP-8 components and load file contents into memory
	components := components.InitComponents(fileName)
	defer components.Display.Destroy()

	running := true

	secondsPerFrame := 1.0 / TARGET_FPS

	for running {
		frameStartTime := time.Now()

		running = handleInput(&components)
		executeInstructions(&components)

		components.DelayTimer.Decrement()
		components.SoundTimer.Decrement()

		if components.SoundTimer.Value != 0 {
			sound.Play(mix.DEFAULT_CHANNELS, 1)
		}

		components.Display.Present()

		elapsedTime := float32(time.Since(frameStartTime).Seconds())

		if elapsedTime < float32(secondsPerFrame) {
			sdl.Delay(uint32((float32(secondsPerFrame) - elapsedTime) * 1000))
		}
	}
}
