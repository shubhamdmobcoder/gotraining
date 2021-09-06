package main

import "fmt"

type Stack []string

func (newStack *Stack) IsEmpty() bool {
	return len(*newStack) == 0
}

func (newStack *Stack) Push(str string) {
	*newStack = append(*newStack, str) 
}


func (newStack *Stack) Pop() (string, bool) {
	if newStack.IsEmpty() {
		return "", false
	} else {
		index := len(*newStack) - 1     
		element := (*newStack)[index]   
		*newStack = (*newStack)[:index] 
		return element, true
	}
}

func main() {
	var stack Stack

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
