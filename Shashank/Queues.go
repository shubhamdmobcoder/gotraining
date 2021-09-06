package main

import (
	"fmt"
)

func enqueue(queue []int, element int) []int {
	queue = append(queue, element) // Appending to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []int) []int {
	element := queue[0] // The first element is the one to be dequeued here
	fmt.Println("Dequeued:", element)
	return queue[1:] // Slicing off the element once it is dequeued.
}

func main() {
	var queue []int // Making a queue of ints.

	queue = enqueue(queue, 5)
	queue = enqueue(queue, 10)
	queue = enqueue(queue, 15)

	fmt.Println("Queue:", queue)

	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	queue = enqueue(queue, 20)
	fmt.Println("Queue:", queue)
}
