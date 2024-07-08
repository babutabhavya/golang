package main

import "fmt"

type StackImp interface {
	push(val int)
	pop() int
	length() int
}

type Stack struct {
	elements []int
}

func (s *Stack) push(values ...int) {
	s.elements = append(s.elements, values...)
}

func (s *Stack) pop() int {
	elementsLength := s.length()
	lastIndex := elementsLength - 1
	if lastIndex >= 0 {
		retval := s.elements[lastIndex]
		s.elements = s.elements[:elementsLength-1]
		return retval
	}
	return -1

}

func (s *Stack) length() int {
	return len(s.elements)
}

func main() {
	var stack Stack = Stack{elements: []int{}}
	stack.push(5)
	fmt.Println(stack.elements)
	fmt.Println(stack.pop())
	fmt.Println(stack.elements)
	fmt.Println(stack.length())
	stack.push(10, 20, 30)
	fmt.Println(stack)
	stack.push(40, 50, 60)
	fmt.Println(stack.length())
	fmt.Println(stack.pop())
}
