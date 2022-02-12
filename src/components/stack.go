package components

import "errors"

//=== Constants ===//

const STACK_SIZE = 16

//=== Struct Definitions ===//

type Stack struct {
	Values  [STACK_SIZE]int16
	Pointer int16
}

//=== Stack Functions ===//

// Creates an empty stack.
func NewStack() *Stack {
	return &Stack{[STACK_SIZE]int16{}, -1}
}

// Pushes the given value onto the stack, unless it is full.
func (stack *Stack) Push(value int16) error {
	if stack.Pointer >= (STACK_SIZE - 1) {
		return errors.New("Stack is full.")
	}

	stack.Pointer++
	stack.Values[stack.Pointer] = value

	return nil
}

// Removes and returns the top value from the stack, unless it is empty.
func (stack *Stack) Pop() (int16, error) {
	if stack.Pointer < 0 {
		return -1, errors.New("Stack is empty.")
	}

	value := stack.Values[stack.Pointer]
	stack.Pointer--

	return value, nil
}
