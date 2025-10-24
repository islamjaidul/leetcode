package main

import (
	"fmt"
)

func findMaxAverageV1(nums []int, k int) float64 {
	var maxAverage float64 = float64(nums[0])
	n := len(nums) - 1
	for i := range nums {
		sum := 0
		j := i
		for j < k+i {
			sum += nums[j]
			j++
			n--
		}
		var average float64 = float64(sum) / float64(k)
		if average > maxAverage {
			maxAverage = average
		}

		if n <= 0 {
			break
		}
	}
	return maxAverage
}

func findMaxAverageV2(nums []int, k int) float64 {
	n := len(nums) - 1
	l := 0
	r := 0
	var totalSum int
	var maxAverage float64 = -1e9

	if n == 0 {
		return float64(nums[0]) / float64(k)
	}

	for r <= n {
		totalSum += nums[r]
		var tempAverage float64
		if r-l == k-1 {
			tempAverage = (float64(totalSum)) / float64(k)
			if tempAverage > maxAverage {
				maxAverage = tempAverage
			}
			totalSum -= nums[l]
			l++
		}

		r++
	}

	return maxAverage

}

func main() {
	// input := []int{1, 12, -5, -6, 50, 3}
	// input := []int{5}
	// input := []int{-1}
	// input := []int{0, 1, 1, 3, 3}
	input := []int{8860, -853, 6534, 4477, -4589, 8646, -6155, -5577, -1656, -5779, -2619, -8604, -1358, -8009, 4983, 7063, 3104, -1560, 4080, 2763, 5616, -2375, 2848, 1394, -7173, -5225, -8244, -809, 8025, -4072, -4391, -9579, 1407, 6700, 2421, -6685, 5481, -1732, -8892, -6645, 3077, 3287, -4149, 8701, -4393, -9070, -1777, 2237, -3253, -506, -4931, -7366, -8132, 5406, -6300, -275, -1908, 67, 3569, 1433, -7262, -437, 8303, 4498, -379, 3054, -6285, 4203, 6908, 4433, 3077, 2288, 9733, -8067, 3007, 9725, 9669, 1362, -2561, -4225, 5442, -9006, -429, 160, -9234, -4444, 3586, -5711, -9506, -79, -4418, -4348, -5891}
	// fmt.Println(findMaxAverageV1(input, 4))

	fmt.Println(findMaxAverageV2(input, 93))
}
