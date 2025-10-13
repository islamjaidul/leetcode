package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}
	hashMap := make(map[int]int)
	for i, num := range nums {
		if remainingValue, exist := hashMap[target-num]; exist && remainingValue != i {
			result = append(result, i, remainingValue)
			break
		}
		hashMap[num] = i
	}
	return result
}

func main() {
	nums := []int{3, 3, 9}
	target := 12
	fmt.Println(twoSum(nums, target))
}
