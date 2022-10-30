package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	word1 = " " + word1
	word2 = " " + word2
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	
	for i := 0; i < m; i ++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				if word1[i] == word2[j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i][j-1] + 1
				}
			} else if i > 0 && j == 0 {
				if word1[i] == word2[j] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] + 1
				}
			} else if i > 0 && j > 0 {
				if word1[i] == word2[j] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
				}
			}
		}
	}
	return dp[m-1][n-1]
}

func min(i, j, k int) int {
	var tmp int
	if i < j {
		tmp = i
	} else {
		tmp = j
	}
	if tmp < k {
		return tmp
	}
	return k
}

func main() {
	fmt.Println(minDistance("horse","ros"))
}