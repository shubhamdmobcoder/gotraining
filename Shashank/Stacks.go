package main

import "fmt"

type Stack []string

// Checking if stack is empty
func (newStack *Stack) IsEmpty() bool {
	return len(*newStack) == 0
}

// Pushing a new value into the stack
func (newStack *Stack) Push(str string) {
	*newStack = append(*newStack, str) //  Appending the new value to the end of the stack
}

// Removing and returning top element of stack. Return false if stack is empty.
func (newStack *Stack) Pop() (string, bool) {
	if newStack.IsEmpty() {
		return "", false
	} else {
		index := len(*newStack) - 1     // Index of the top most element.
		element := (*newStack)[index]   // Index into the slice and obtain the element.
		*newStack = (*newStack)[:index] // Removing it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	var stack Stack // Creating a stack variable of type Stack

	stack.Push("hi")
	stack.Push("say")
	stack.Push("Shashank")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
