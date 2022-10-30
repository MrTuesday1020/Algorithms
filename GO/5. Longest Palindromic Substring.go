package main

import "fmt"

// 动态规划
//func longestPalindrome(s string) string {
//	l := len(s)
//	maxLen, start, end := 0, 0, 0
//	var dp [][]int
//	
//	for i := 0; i < l + 1; i++ {
//		arr := make([]int, l + 1)
//		dp = append(dp, arr)
//	}
//	
//	for i := 0; i < l; i++ {
//		for j := 0; j < l ; j++ {
//			if s[i] == s[l-j-1] {
//				dp[i+1][j+1] = dp[i][j] + 1
//				if dp[i+1][j+1] > maxLen {
//					beforeReverseID := l - j - 1    					// 当前最后一位的下标
//					lastSubStrID := beforeReverseID + dp[i+1][j+1] - 1	// 这里不好理解
//					if lastSubStrID == i {
//						maxLen = dp[i+1][j+1]
//						start = i+1-maxLen
//						end = i+1
//					}
//				}
//			}
//		} 
//	}
//	
//	for i := 0; i < l + 1; i++ {
//		fmt.Println(dp[i])
//	}
//	fmt.Println(start, end)
//		
//	return s[start:end]
//}

// 中心扩展法 便于理解
func longestPalindrome(s string) string {
	maxLen := 0
	retStr := ""
	for i := 0; i < len(s); i++ {
		len1, str1 := expandFromCenter(s, i, i)
		len2, str2 := expandFromCenter(s, i, i + 1)
		if len1 > maxLen {
			maxLen = len1
			retStr = str1
		}
		if len2 > maxLen {
			maxLen = len2
			retStr = str2
		}
	}
	
	return retStr
}

func expandFromCenter(s string, left, right int) (int, string) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return right - left - 1, s[left+1:right]
}

func main() {
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aacabdkacaa"))
}