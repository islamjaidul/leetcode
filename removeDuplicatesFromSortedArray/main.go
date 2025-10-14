package main

import "fmt"

func removeDuplicatesV1(nums []int) int {
	i := 0
	temp := -101

	for _, num := range nums {
		if num > temp {
			temp = num
			nums[i] = num
			i++
		}
	}

	return i

}

func main() {
	input := []int{-100, -99, -98, -98, -97, -97}
	fmt.Println(removeDuplicatesV1(input))
	fmt.Println(input)

}
