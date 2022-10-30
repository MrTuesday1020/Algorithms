package main

import "fmt"

// https://leetcode.cn/problems/regular-expression-matching/solution/shi-pin-tu-jie-dong-tai-gui-hua-zheng-ze-biao-da-s/
func isMatch(s string, p string) bool {
	ls := len(s)
	lp := len(p)
	var dp [][]bool
	for i := 0; i < ls + 1; i++{
		dp = append(dp, make([]bool, lp+1))
		if i == 0 {
			dp[i][0] = true
		}
	}	
	
	for j := 0; j < lp; j ++ {
		if p[j] == '*' {
			dp[0][j+1] = dp[0][j+1-2]
		}	
	}
	
	for i := 0; i < ls+1; i++ {
		fmt.Println(dp[i])
	}
	fmt.Println("=========")
	
	for i := 0; i < ls; i++ {
		for j := 0; j < lp; j++ {
			if s[i] == p[j] || p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				if dp[i+1][j+1-2] {
					dp[i+1][j+1] = true	
				} else if s[i] == p[j-1] || p[j-1] == '.' {
					dp[i+1][j+1] = dp[i][j+1]
				}
			}
		}
	}
	
	for i := 0; i < ls+1; i++ {
		fmt.Println(dp[i])
	}
	
	return dp[ls][lp]
}

func main() {
	//fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aaa", "ab*a*c*a"))
	//fmt.Println(isMatch("ssippi", "s*p*."))
}