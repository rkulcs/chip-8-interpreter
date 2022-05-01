package instructions

import (
	"github.com/veandco/go-sdl2/sdl"
)

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

func GetInputKeyValue(keyCode int) (code byte, pressed bool) {

	if keyCode == 0 {
		return 0x0, false
	}

	switch keyCode {
	case KEY_0:
		return 0x0, true
	case KEY_1:
		return 0x1, true
	case KEY_2:
		return 0x2, true
	case KEY_3:
		return 0x3, true
	case KEY_4:
		return 0x4, true
	case KEY_5:
		return 0x5, true
	case KEY_6:
		return 0x6, true
	case KEY_7:
		return 0x7, true
	case KEY_8:
		return 0x8, true
	case KEY_9:
		return 0x9, true
	case KEY_A:
		return 0xA, true
	case KEY_B:
		return 0xB, true
	case KEY_C:
		return 0xC, true
	case KEY_D:
		return 0xD, true
	case KEY_E:
		return 0xE, true
	case KEY_F:
		return 0xF, true
	}

	return 0x0, false
}
