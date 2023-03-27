# CHIP-8 Interpreter

A CHIP-8 interpreter written in Go.

# Acknowledgements

The following references were used to implement the components and instruction set of CHIP-8:

- https://tobiasvl.github.io/blog/write-a-chip-8-emulator
- http://devernay.free.fr/hacks/chip8/C8TECH10.HTM

The sound effect used by the interpreter was generated using [jsfxr](https://sfxr.me/).

# Requirements

- [SDL2 binding for Go](https://github.com/veandco/go-sdl2)

# Usage

```
cd src/main
go build .
./main
```

# Screenshots

Pong:<br><br>
<img src="https://user-images.githubusercontent.com/50153954/227824820-915958f0-b5c1-440a-a8d6-05bd8f222eaa.png" width="400" /><br><br>
Tetris:<br><br>
<img src="https://user-images.githubusercontent.com/50153954/227824824-c6d36ec4-10b9-431f-87e2-bbb96febd185.png" width="400" /><br><br>
Space Invaders:<br><br>
<img src="https://user-images.githubusercontent.com/50153954/227824826-2e07ce51-82ce-49f7-bc66-f219a1adeb6c.png" width="400" /><br><br>
