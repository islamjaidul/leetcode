package main

import "fmt"

func missingMultiple(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	hashMap := make(map[int]int)
	for _, value := range nums {
		hashMap[value]++
	}

	for i := 1; ; i++ {
		multiple := i * k
		if _, exist := hashMap[multiple]; !exist {
			return multiple
		}
	}
}

func main() {
	input := []int{1, 4, 7, 10, 15}
	fmt.Println(missingMultiple(input, 5))
}
