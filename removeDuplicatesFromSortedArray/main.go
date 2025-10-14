package main

import "fmt"

func removeDuplicatesV1(nums []int) int {
	// resultArr := make([]int, len(nums))
	i := 0
	temp := nums[0]

	for num := range nums {
		if num > temp {
			nums[i] = num
			i++
		}
	}

	return i

}

func main() {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicatesV1(input))

}
