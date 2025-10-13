package main

import (
	"fmt"
)

func sortedSquaresV1(nums []int) []int {
	negativeValue := []int{}
	positiveValue := []int{}

	for _, num := range nums {
		square := num * num
		if num < 0 {
			negativeValue = append([]int{square}, negativeValue...)
		} else {
			positiveValue = append(positiveValue, square)
		}
	}

	if len(negativeValue) == 0 {
		return positiveValue
	}

	if len(positiveValue) == 0 {
		return negativeValue
	}

	i := len(negativeValue) - 1
	j := len(positiveValue) - 1
	k := len(nums) - 1

	for j >= 0 && i >= 0 {
		if negativeValue[i] > positiveValue[j] {
			nums[k] = negativeValue[i]
			i--
		} else {
			nums[k] = positiveValue[j]
			j--
		}
		k--
	}

	// Add remaining elements (if any)
	for i >= 0 {
		nums[k] = negativeValue[i]
		i--
		k--
	}
	for j >= 0 {
		nums[k] = positiveValue[j]
		j--
		k--
	}

	return nums

}

func main() {
	input := []int{-1}
	fmt.Println(sortedSquaresV1(input))
}
