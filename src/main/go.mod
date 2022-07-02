module main

go 1.17

replace components => ../components

require (
	components v0.0.0-00010101000000-000000000000
	instructions v0.0.0-00010101000000-000000000000
)

require github.com/veandco/go-sdl2 v0.4.24

replace instructions => ../instructions
