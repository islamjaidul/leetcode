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

func main() {
	input := []int{3, 2, 3, 2}
	fmt.Println(majorityElementV1(input))
}
