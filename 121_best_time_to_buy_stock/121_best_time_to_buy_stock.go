package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	result := 0

	i := 0
	j := len(prices) - 1

	for i < j {
		diff := prices[j] - prices[i]
		fmt.Printf("%d - %d = %d\n", prices[j], prices[i], diff)
		if diff < 0 {
			i++
		} else {
			if diff > result {
				result = diff
			} else {
				j--
			}
			i++
		}

	}

	return result
}

func main() {
	prices := [][]int{
		{2, 1, 4},
		{7, 1, 5, 3, 6, 4},
	}

	for _, priceArr := range prices {
		fmt.Printf("Max Profit can be made is: %d\n", maxProfit(priceArr))
	}
}
