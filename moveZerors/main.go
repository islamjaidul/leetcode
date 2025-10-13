package main

import "fmt"

func moveZeroes(nums []int) {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	input := []int{0, 0, 1}
	moveZeroes(input)
	fmt.Println(input)
}
