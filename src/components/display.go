package components

import (
	"github.com/veandco/go-sdl2/sdl"
)

//=== Constants ===//

var Font [][]byte = [][]byte{{0xF0, 0x90, 0x90, 0x90, 0xF0},
	{0x20, 0x60, 0x20, 0x20, 0x70},
	{0xF0, 0x10, 0xF0, 0x80, 0xF0},
	{0xF0, 0x10, 0xF0, 0x10, 0xF0},
	{0x90, 0x90, 0xF0, 0x10, 0x10},
	{0xF0, 0x80, 0xF0, 0x10, 0xF0},
	{0xF0, 0x80, 0xF0, 0x90, 0xF0},
	{0xF0, 0x10, 0x20, 0x40, 0x40},
	{0xF0, 0x90, 0xF0, 0x90, 0xF0},
	{0xF0, 0x90, 0xF0, 0x10, 0xF0},
	{0xF0, 0x90, 0xF0, 0x90, 0x90},
	{0xE0, 0x90, 0xE0, 0x90, 0xE0},
	{0xF0, 0x80, 0x80, 0x80, 0xF0},
	{0xE0, 0x90, 0x90, 0x90, 0xE0},
	{0xF0, 0x80, 0xF0, 0x80, 0xF0},
	{0xF0, 0x80, 0xF0, 0x80, 0x80}}

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
	pixels  [DISPLAY_HEIGHT][DISPLAY_WIDTH]bool
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

	return Display{window, surface, [DISPLAY_HEIGHT][DISPLAY_WIDTH]bool{}}
}

// Destroys the SDL window of the display.
func (display *Display) Destroy() {
	display.window.Destroy()
}

// Draws a pixel at the specified coordinates. If the value of "on" is true,
// then a white pixel will be drawn if there is a black pixel at the specified
// location. Otherwise, a black pixel will be drawn. True is returned if the
// pixel at the given coordinates was white, otherwise, false is returned.
func (display *Display) Draw(x int32, y int32, on bool) (wasOn bool) {
	rect := sdl.Rect{x * DISPLAY_SCALE, y * DISPLAY_SCALE, DISPLAY_SCALE, DISPLAY_SCALE}

	wasOn = display.pixels[y][x]

	if wasOn && on {
		display.pixels[y][x] = false
		display.surface.FillRect(&rect, COLOR_BLACK)
	} else if !wasOn && on {
		display.pixels[y][x] = true
		display.surface.FillRect(&rect, COLOR_WHITE)
	}

	display.window.UpdateSurface()
	return
}

// Clears the display.
func (display *Display) Clear() {
	display.surface.FillRect(nil, COLOR_BLACK)
	display.window.UpdateSurface()
}

// Gets the location of the font of the provided hexadecimal digit.
func (display *Display) GetFontLocation(digit byte) int {
	return FONT_START_LOCATION + (5 * int(digit))
}
