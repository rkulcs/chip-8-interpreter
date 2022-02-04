package components

//=== Struct Definitions ===//

type Components struct {
	Display   *Display
	Registers *Registers
	Memory    *Memory
}

// Component Functions ===//

func InitComponents(fileName string) Components {
	display := InitDisplay()
	registers := &Registers{}
	memory := InitMemory(fileName)

	return Components{&display, registers, &memory}
}
