package components

//=== Struct Definitions ===//

type Components struct {
	Display    *Display
	Registers  *Registers
	Memory     *Memory
	Stack      *Stack
	DelayTimer *Timer
	SoundTimer *Timer
	InputMap   *InputStates
}

// Component Functions ===//

func InitComponents(fileName string) Components {
	display := InitDisplay()
	registers := &Registers{}
	memory := InitMemory(fileName)
	stack := NewStack()
	delayTimer := &Timer{}
	soundTimer := &Timer{}
	inputMap := &InputStates{}

	registers.PC = 0x200

	return Components{&display, registers, &memory, stack, delayTimer, soundTimer, inputMap}
}
