package main

import (
	"fmt"
)

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {

	intSliceValues := []int{1, 5, 5, 6, 8, 4, 5, 3, 5, 1, 1, 5, 6}
	fmt.Println(intSliceValues)
	removeDuplicateValuesSlice := removeDuplicateValues(intSliceValues)
	fmt.Println(removeDuplicateValuesSlice)
}
