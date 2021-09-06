package main

import "fmt"

func main() {
	a := []int{100, 200, 300, 100, 200, 400, 500, 55, 0}
	fmt.Println(a)

	x := remdup(a)
	fmt.Println(x)
}

func remdup(a []int) []int {
	y := map[int]bool{}
	x := []int{}

	for v := range a {
		if y[a[v]] == true {
		} else {
			y[a[v]] = true
			x = append(x, a[v])
		}
	}
	return x
}
