package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for i := range prices {
		if prices[i] < minPrice && prices[i] > 0 {
			minPrice = prices[i]
		}
		transaction := prices[i] - minPrice
		if transaction > profit {
			profit = transaction
		}
	}

	return profit

}

func main() {
	input := []int{7, 2, 5, 0, 6, 4, 1, 9}
	fmt.Println(maxProfit(input))
}
