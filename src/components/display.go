package components

import "github.com/veandco/go-sdl2/sdl"

//=== Constants ===//

// The value by which the display should be scaled.
const DISPLAY_SCALE = 10

// The original width of the display.
const DISPLAY_WIDTH = 64

// The original height of the display.
const DISPLAY_HEIGHT = 32

const COLOR_WHITE = 0xffffffff
const COLOR_BLACK = 0x00000000

//=== Struct Definitions ===//

type Display struct {
	window  *sdl.Window
	surface *sdl.Surface
}

//=== Display Functions ===//

// Creates and returns a new display.
func InitDisplay() Display {
	window, err := sdl.CreateWindow("CHIP-8 Interpreter",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		DISPLAY_WIDTH*DISPLAY_SCALE,
		DISPLAY_HEIGHT*DISPLAY_SCALE, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	surface, err := window.GetSurface()

	if err != nil {
		panic(err)
	}

	return Display{window, surface}
}

// Destroys the SDL window of the display.
func (display *Display) Destroy() {
	display.window.Destroy()
}

// Draws a white pixel at the specified coordinates.
func (display *Display) Draw(x int32, y int32) {
	rect := sdl.Rect{x, y, DISPLAY_SCALE, DISPLAY_SCALE}
	display.surface.FillRect(&rect, COLOR_WHITE)
	display.window.UpdateSurface()
}
