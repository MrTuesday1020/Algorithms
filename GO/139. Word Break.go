package main

import "fmt"

// do recursion, easy to understand, but out of time
//var ans bool
//
//func wordBreak(s string, wordDict []string) bool {
//	ans = false
//	recursion(s, "", wordDict)
//	return ans
//}
//
//func recursion(target, current string, wordDict []string) {
//	if target == current {
//		ans = ans || true
//		return
//	}
//	if len(current) >= len(target) {
//		ans = ans || false
//		return
//	}
//	if current != target[:len(current)] {
//		ans = ans || false
//		return
//	}
//	for _, str := range wordDict {
//		recursion(target, current + str, wordDict)
//	}
//}


// dp
// '' l e e t c o d e      	target
//  t f f f t f f f t		dp table
func wordBreak(s string, wordDict []string) bool {
	// build a map
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	
	// build dp tabel
	dp := make([]bool, len(s)+1)
	dp[0] = true
	
	// compute dp table
	for i := 0; i < len(s); i++ {
		for j := i+1; j < len(s)+1; j++ {
			if dp[i] && m[s[i:j]] {
				dp[j] = true
			}
		}
	}
	
	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}