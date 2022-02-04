package main

import (
	"bufio"
	"components"
	"fmt"
	"os"

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
	// Get the file name of the CHIP-8 program to run.
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
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}
}
