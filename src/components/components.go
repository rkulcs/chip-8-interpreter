package components

//=== Struct Definitions ===//

type Components struct {
	Display   *Display
	Registers *Registers
	Memory    *Memory
	Stack     *Stack
}

// Component Functions ===//

func InitComponents(fileName string) Components {
	display := InitDisplay()
	registers := &Registers{}
	memory := InitMemory(fileName)
	stack := NewStack()

	registers.PC = 0x200

	return Components{&display, registers, &memory, stack}
}
