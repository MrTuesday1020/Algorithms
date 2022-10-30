//You are given two strings s1 and s2 of equal length. A string swap is an 
//operation where you choose two indices in a string (not necessarily different) and 
//swap the characters at these indices. 
//
// Return true if it is possible to make both strings equal by performing at 
//most one string swap on exactly one of the strings. Otherwise, return false. 
//
// 
// Example 1: 
//
// 
//Input: s1 = "bank", s2 = "kanb"
//Output: true
//Explanation: For example, swap the first character with the last character of 
//s2 to make "bank".
// 
//
// Example 2: 
//
// 
//Input: s1 = "attack", s2 = "defend"
//Output: false
//Explanation: It is impossible to make them equal with one string swap.
// 
//
// Example 3: 
//
// 
//Input: s1 = "kelb", s2 = "kelb"
//Output: true
//Explanation: The two strings are already equal, so no string swap operation 
//is required.
// 
//
// 
// Constraints: 
//
// 
// 1 <= s1.length, s2.length <= 100 
// s1.length == s2.length 
// s1 and s2 consist of only lowercase English letters. 
// 
//
// Related Topics 哈希表 字符串 计数 👍 63 👎 0

package main

import "fmt"

func minSwap(nums1 []int, nums2 []int) int {
	var dp = make([][]int, len(nums1))
	for i := range nums1 {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < len(nums1); i++ {
		// 可换可不换
		if (nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i]) && (nums1[i-1] < nums2[i] && nums2[i-1] < nums1[i]) {
			dp[i][0] = min(dp[i-1][0], dp[i-1][1])
			dp[i][1] = dp[i][0] + 1
		} else if nums1[i-1] < nums1[i] && nums2[i-1] < nums2[i] {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		} else {
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		}
	}
	return min(dp[len(nums1)-1][0], dp[len(nums1)-1][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minSwap([]int{1,3,5,4},[]int{1,2,3,7}))
}