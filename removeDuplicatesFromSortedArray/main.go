package main

import "fmt"

func removeDuplicatesV1(nums []int) int {
	i := 0
	// -100 <= nums[i] <= 100 was the constrain, that's why I initiated with -101
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

func removeDuplicatesV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 1 // first item will be sorted and not duplicate always
	// j = 1 started due to compare nums[j] != nums[j-1
	for j := 1; j < len(nums); j++ {
		// compare with current nums[j] and previous nums[j-1]
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}

	// we don't need to do anything with the last few duplicate items as problem statement said that, it will assert the count and sorted item

	return i

}

func main() {
	input := []int{-100, -99, -98, -98, -97, -97}
	// fmt.Println(removeDuplicatesV1(input))
	fmt.Println(removeDuplicatesV2(input))
	fmt.Println(input)

}
