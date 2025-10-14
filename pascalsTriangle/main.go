package main

import "fmt"

func generate(numRows int) [][]int {
	resultArr := make([][]int, numRows)
	for i := range numRows {
		triangleArr := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				triangleArr[j] = 1
			} else {
				triangleArr[j] = resultArr[i-1][j-1] + resultArr[i-1][j]
			}

		}
		resultArr[i] = triangleArr
	}

	// [1] [1 1] [1 0 1] [1 0 0 1] [1 0 0 0 1]]
	// [1] [1 1] [1 2 1] [1 2 3 1] [1 2 3 4 1]]
	// [1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1], [1,5,10,10,5,1]

	return resultArr
}

func main() {
	input := 5
	fmt.Println(generate(input))
}
