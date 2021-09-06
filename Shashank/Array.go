package twoSum

func twoSum(array []int, target int) []int {
	for i := 0; i < len(array); i++ {
		var start = array[i]
		for j := i + 1; j < len(array); j++ {
			var end = array[j]
			if (start + end) == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
