package main

import "fmt"

// This is much faster
// Only one forward pass (O(n))
// No unnecessary swapping
// No nested conditions
func moveZeroesV1(nums []int) {
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

func moveZeroesV2(nums []int) {
	n := len(nums) - 1
	left := 0
	right := 0

	for left <= n && right <= n {
		if nums[left] != 0 && nums[right] != 0 {
			right++
			left++
		} else if nums[left] == 0 && nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		} else {
			right++
		}
	}
}

func moveZeroesV3(nums []int) {
	left := 0
	for right := range nums {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}

func main() {
	input := []int{1, 0, 12, 23, 0, 0, 25}
	// input := []int{-1, 0, 0, 1, 0}
	moveZeroesV2(input)
	fmt.Println(input)
}
