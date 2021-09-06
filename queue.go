package main

import "fmt"

func main(){
	var queue[] int 
	queue = enqueue(queue , 1)
	queue = enqueue(queue , 20)
	queue = enqueue(queue , 3)
	fmt.Println(queue)

	var dqueue[] int
	dqueue = dequeue(queue)
	fmt.Println("Queue",dqueue)


}

func enqueue(queue[] int , x int) []int {
	queue = append(queue, x )
	fmt.Println("enquied : ", x)
	return queue
}

func dequeue(queue [] int)[]int{
	element := queue[0]
	fmt.Println("dequeue:",element)
	return queue[1:]

}
