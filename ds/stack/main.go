package main

import "fmt"

const StackSize = 1024

// Stack - first in last out
type Stack struct {
	stack    []string
	stackptr int // Always point to the next free slot. Top of the stack is stack[sp-1]
	// Incremented and decremented as the stack grows or shrinks.
}

func New() *Stack {
	return &Stack{stack: make([]string, StackSize)}
}

func (s *Stack) Push(data string) error {
	if s.stackptr >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	s.stack[s.stackptr] = data
	s.stackptr++
	return nil
}

func (s *Stack) Pop() string {
	data := s.stack[s.stackptr-1]
	s.stackptr--
	return data
}

func (s *Stack) Size() int {
	return s.stackptr
}

func (s *Stack) Print() {
	for i, data := range s.stack {
		if i >= s.stackptr {
			break
		}
		fmt.Println(data)
	}
}

func main() {
	myStack := New()
	myStack.Push("GoodBye")
	myStack.Push("Hello")
	myStack.Push("World")
	myStack.Print()
	fmt.Println("Stack size: ", myStack.Size())

	fmt.Println("Pop: ", myStack.Pop())
	fmt.Println("Stack size: ", myStack.Size())
	fmt.Println("Pop: ", myStack.Pop())
	fmt.Println("Stack size: ", myStack.Size())
	fmt.Println("Pop: ", myStack.Pop())
	fmt.Println("Stack size: ", myStack.Size())
}
