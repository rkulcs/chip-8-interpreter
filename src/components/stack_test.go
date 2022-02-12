package components

import "testing"

func TestStack(test *testing.T) {
	stack := NewStack()

	var i int16

	for i = 0; i < 16; i++ {
		stack.Push(i)

		if stack.Pointer != i {
			test.Errorf("Incorrect stack pointer value.")
		}
	}

	for i = 15; i >= 0; i-- {
		value, err := stack.Pop()

		if err != nil {
			test.Errorf("An error occurred when popping from the stack.")
		}

		if stack.Pointer != (i - 1) {
			test.Errorf("Incorrect stack pointer value.")
		}

		if value != i {
			test.Errorf("Incorrect stack value.")
		}
	}
}
