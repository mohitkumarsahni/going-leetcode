package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minBuyPrice := math.MaxInt64
	maxProfit := 0

	for _, price := range prices {
		if price < minBuyPrice {
			minBuyPrice = price
		} else if (price - minBuyPrice) > maxProfit {
			maxProfit = price - minBuyPrice
		}
	}

	return maxProfit
}

func main() {
	slice := []int{2,0,2}

	fmt.Println("Most Profit : ",maxProfit(slice))
}