package main

import (
	"fmt"
)

// dp[i] = 1 + min{dp[i-coins[k]]}
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	
	for i := 1; i < amount + 1; i++ {
		dp[i] = amount+1
	} 
	
	fmt.Println(dp)
	
	for i := 1; i < amount + 1; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i - coins[j]] + 1);
			}
		}
	}
	
	fmt.Println(dp)
	
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(coinChange([]int{2,7},10))
}