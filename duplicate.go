package main

import "fmt"

func main() {
	a := []int{100, 200, 300, 100, 200, 400, 500, 55, 0}
	fmt.Println(a)

	// test for method
	x := remdup(a)
	fmt.Println(x)
}

func remdup(a []int) []int {
	//use map to record duplicates
	y := map[int]bool{}
	x := []int{}

	for v := range a {
		if y[a[v]] == true {
			// here we don't add duplicates
		} else {
			y[a[v]] = true
			//store or append the data into x
			x = append(x, a[v])
		}
	}
	return x
}
