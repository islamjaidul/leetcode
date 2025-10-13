package main

import "fmt"

func pivotIndex(nums []int) int {
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	length := len(nums) - 1
	sum := 0
	pivotIndex := -1

	// Calculate left running sum
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		leftSum[i] = sum
	}

	// Calculate right running sum
	sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[length]
		rightSum[length] = sum
		length--
	}

	// Find pivot where left sum equals right sum
	for i := 0; i < len(nums); i++ {
		if leftSum[i] == rightSum[i] {
			pivotIndex = i
			break
		}
	}

	return pivotIndex
}

func main() {
	input := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(input))
}
