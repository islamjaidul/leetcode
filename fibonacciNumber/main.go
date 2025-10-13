package main

import "fmt"

var dp = make(map[int]int)

func fib(n int) int {

	if n <= 1 {
		return n
	}

	if value, isExist := dp[n]; isExist {
		return value
	}

	dp[n] = fib(n-1) + fib(n-2)
	return dp[n]
}

func main() {
	input := 50
	fmt.Println(fib(input))
}
