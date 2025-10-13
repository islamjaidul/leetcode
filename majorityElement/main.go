package main

import "fmt"

func majorityElementV1(nums []int) int {
	hashMap := make(map[int]int)
	for _, num := range nums {
		if value, isExist := hashMap[num]; isExist != false {
			hashMap[num] = value + 1
		} else {
			hashMap[num] = 1
		}
	}

	maxValue := 0
	result := 0
	for key, value := range hashMap {
		if value >= maxValue {
			maxValue = value
			result = key
		}
	}
	return result
}

func majorityElementV2(nums []int) int {
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
		if hashMap[num] > len(nums)/2 {
			return num
		}
	}
	return nums[0]
}

/*
 * Explanation
 * We are selecting candidate based on the vote
 * Each time we are checking and changing the vote
 * We are returning the last candidate for this
 * Example - [3, 2, 3, 2, 2] for this voting is [candidate = vote] = [3 = 1, 2 = 0, 3 = 1, 2 = 0, 2 = 1]
 * For the above example - 2 was the last candidate that we returned
 */

func majorityElementV3(nums []int) int {
	vote := 0
	candidate := 0

	for _, num := range nums {
		if vote == 0 {
			candidate = num
		}
		if num == candidate {
			vote++
		} else {
			vote--
		}
		fmt.Println(vote)
	}

	return candidate
}

func main() {
	input := []int{3, 2, 3, 2, 2}
	// fmt.Println(majorityElementV1(input))
	// fmt.Println(majorityElementV2(input))
	fmt.Println(majorityElementV3(input))
}
