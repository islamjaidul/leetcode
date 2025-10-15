package main

import (
	"fmt"
	"sort"
)

func mergeV1(intervals [][]int) [][]int {

	// Doing sort to make sure we can check overlap correctly with [i-1] and [i]
	// resultArr where we hold the merged interval and cross check with intervals each iteration
	// If any overlap, then merge it in resultArr

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	resultArr := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := resultArr[len(resultArr)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			last[1] = max(last[1], current[1])
			resultArr[len(resultArr)-1] = last
		} else {
			resultArr = append(resultArr, current)
		}
	}

	return resultArr
}

func main() {
	// [[4,7],[1,4]]
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {11, 18}}
	// input := [][]int{{4, 7}, {1, 4}}
	fmt.Println(mergeV1(input))
}
