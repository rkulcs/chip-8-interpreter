package components

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
)

//=== Struct Definitions ===//

type InputStates struct {
	isPressed [16]bool
}

const KEY_1 = sdl.K_1
const KEY_2 = sdl.K_2
const KEY_3 = sdl.K_3
const KEY_C = sdl.K_4

const KEY_4 = sdl.K_q
const KEY_5 = sdl.K_w
const KEY_6 = sdl.K_e
const KEY_D = sdl.K_r

const KEY_7 = sdl.K_a
const KEY_8 = sdl.K_s
const KEY_9 = sdl.K_d
const KEY_E = sdl.K_f

const KEY_A = sdl.K_z
const KEY_0 = sdl.K_x
const KEY_B = sdl.K_c
const KEY_F = sdl.K_v

func (inputStates *InputStates) SetInputKeyState(keyCode int, newState bool) {
	index, err := GetInputKeyValue(keyCode)

	if err != nil {
		return
	}

	inputStates.isPressed[index] = newState
}

func (inputStates *InputStates) GetInputKeyState(code byte) bool {
	return inputStates.isPressed[code]
}

func (inputStates *InputStates) IsAnyKeyPressed() (bool, byte) {
	var i byte

	for i = 0x0; i <= 0xF; i++ {
		if inputStates.GetInputKeyState(i) {
			return true, i
		}
	}

	return false, 0
}

func GetInputKeyValue(keyCode int) (code byte, err error) {
	switch keyCode {
	case KEY_0:
		code = 0x0
	case KEY_1:
		code = 0x1
	case KEY_2:
		code = 0x2
	case KEY_3:
		code = 0x3
	case KEY_4:
		code = 0x4
	case KEY_5:
		code = 0x5
	case KEY_6:
		code = 0x6
	case KEY_7:
		code = 0x7
	case KEY_8:
		code = 0x8
	case KEY_9:
		code = 0x9
	case KEY_A:
		code = 0xA
	case KEY_B:
		code = 0xB
	case KEY_C:
		code = 0xC
	case KEY_D:
		code = 0xD
	case KEY_E:
		code = 0xE
	case KEY_F:
		code = 0xF
	default:
		err = errors.New("invalid key code")
	}

	return
}
