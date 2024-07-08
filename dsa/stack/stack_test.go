package main

import (
	"testing"
)

func setupStack() *Stack {
	return &Stack{elements: []int{}}
}
func TestStac(t *testing.T) {

	t.Run("Initialise an empty Stack", func(t *testing.T) {
		stack := setupStack()
		if stack.length() != 0 {
			t.Error("Invalid error while initialise an empty stack")
		}
	})
	t.Run("Push increasing the Stack size by 1", func(t *testing.T) {
		stack := setupStack()
		stack.push(1)
		if stack.length() != 1 {
			t.Error("Invalid error while initialise an empty stack")
		}
	})

	t.Run("Popping decreasing the Stack size by 1", func(t *testing.T) {
		stack := setupStack()
		stack.push(1)

		if stack.length() != 1 {
			t.Error("Invalid error while initialise an empty stack")
		}
		var retval int = stack.pop()
		if retval != 1 {
			t.Error("Pop returned an unexpected value")
		}
	})

	t.Run("Popping when the stack size is already zero", func(t *testing.T) {
		stack := setupStack()
		var retval int = stack.pop()
		if retval != -1 {
			t.Error("Pop returned an unexpected value")
		}
	})
}
