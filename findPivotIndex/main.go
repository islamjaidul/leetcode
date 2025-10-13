package main

import "fmt"

func pivotIndexV1(nums []int) int {
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	length := len(nums) - 1
	sum := 0
	pivotIndex := -1

	for i := range nums {
		sum += nums[i]
		leftSum[i] = sum
	}

	sum = 0
	for range nums {
		sum += nums[length]
		rightSum[length] = sum
		length--
	}

	for i := range nums {
		if leftSum[i] == rightSum[i] {
			pivotIndex = i
			break
		}
	}

	return pivotIndex
}

func pivotIndexV2(nums []int) int {
	total := 0
	leftSum := 0

	// Calculate total sum first
	for _, num := range nums {
		total += num
	}

	// Check each index
	for i := range nums {
		total -= nums[i]
		if leftSum == total {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}

func main() {
	input := []int{1, 7, 3, 6, 5, 6}
	// fmt.Println(pivotIndexV1(input))
	fmt.Println(pivotIndexV2(input))
}
