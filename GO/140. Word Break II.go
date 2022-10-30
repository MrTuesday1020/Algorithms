package main

import "fmt"

var ans []string

func wordBreak(s string, wordDict []string) []string {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
		
	ans = []string{}
	search(s, m, "")
	return ans
}

// if we got a matched prefix, then we add prefx to result, and do dfs on suffix
func search(s string, m map[string]bool, sub string) {
	if s == "" {
		ans = append(ans, sub)
		return 
	}
	for j := 1; j < len(s)+1; j++ {
		if m[s[:j]] {
			if sub == "" {
				search(s[j:], m, s[:j])
			} else {
				search(s[j:], m, sub + " " + s[:j])
			}
		}
	}
}

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat","cats","and","sand","dog"}))
}