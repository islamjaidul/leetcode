package main

import "fmt"

func longestSubArraySum(nums []int, k int) int {
	n := len(nums) - 1
	l := 0
	r := 0
	result := 0
	sum := 0

	for r <= n {
		sum += nums[r]
		if sum < k {
			result = sum
		} else {
			sum -= nums[l]
			l++
		}
		r++
	}

	return result
}

func longestSubArray(nums []int, k int) []int {
	n := len(nums) - 1
	l := 0
	r := 0
	hashMap := make(map[int]int)
	sum := 0
	result := []int{}

	for r <= n {
		sum += nums[r]
		hashMap[r] = nums[r]
		if sum < k {
			delete(hashMap, l)
			l++
		}
		r++
	}

	for _, value := range hashMap {
		result = append(result, value)
	}

	return result
}

func main() {
	input := []int{1, 2, 3}
	k := 3
	fmt.Println(longestSubArraySum(input, k))
}
