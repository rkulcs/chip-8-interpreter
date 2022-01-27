package components

import (
	"testing"
)

func TestMemoryAccess(test *testing.T) {
	memory := Memory{[MEM_SIZE]byte{}}

	indices := [...]int{0, 32, 54, 100, MEM_SIZE - 1}

	for _, index := range indices {
		_, err := memory.ReadFrom(index)

		if err != nil {
			test.Errorf("Cannot read memory contents at address %v", index)
		}
	}
}

func TestMemoryWriting(test *testing.T) {
	memory := Memory{[MEM_SIZE]byte{}}

	indices := [...]int{0, 2, 65, 879, MEM_SIZE - 1}
	newValues := [...]byte{0, 200, 3, 87, 255}

	for i, address := range indices {
		err := memory.WriteTo(address, newValues[i])

		if err != nil {
			test.Errorf("Could not write to memory at address %v", address)
		}

		value, _ := memory.ReadFrom(address)

		if value != newValues[i] {
			test.Errorf("The value at address %v was not modified", address)
		}
	}
}
