package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	resultArr := [][]int{}
	slices.Sort(nums)

	for i, _ := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				resultArr = append(resultArr, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[j] == nums[k+1] {
					k--
				}
			}
		}
	}

	return resultArr
}

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	// [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum(input))
}
