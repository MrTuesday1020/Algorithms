package main

import "fmt"

func maxProfit(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
		}
	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}