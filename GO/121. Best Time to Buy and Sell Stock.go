package main

import "fmt"

func maxProfit(prices []int) int {	
	minPrice := prices[0]
	maxProfit := 0
	
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			if prices[i] - minPrice > maxProfit {
				maxProfit = prices[i] - minPrice
			}
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,5,3}))
}