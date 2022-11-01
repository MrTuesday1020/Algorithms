package main

import "fmt"

func isMatch(s string, p string) bool {
	s = " " + s
	p = " " + p
	m := len(s)
	n := len(p)
	
	// init dp tabale
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	
	// compute dp table
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {	// init left top conner, " " == " "
				dp[i][j] = true
			} else if j == 0 { 		// init first column
				dp[i][j] = false
			} else if i == 0 {		// init first row
				if p[j] == '*' {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = false
				}
			} else {				// compute other cells
				if p[j] == '*' {
					dp[i][j] = dp[i][j-1] || dp[i-1][j]
				} else if p[j] == '?' || s[i] == p[j]{
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = false
				}
			}
		}
	}
	
	return dp[m-1][n-1]
}

func main() {
	
}