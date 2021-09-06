package main

import "fmt"

func main() {

	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size) // taking input(size of array) from user

	var arr = make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
	}

	fmt.Printf("Enter the desired output:  ")
	var target int
	fmt.Scanln(&target)

	fmt.Println("Your array is: ", arr)

	// iterate for adding element of array
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				print("[", i, ", ", j, "]\n")
			} else {
				fmt.Println("target is not valid")
			}

		}

	}
}
