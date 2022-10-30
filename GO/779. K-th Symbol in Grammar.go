package main

//We build a table of n rows (1-indexed). We start by writing 0 in the 1ˢᵗ row. 
//Now in every subsequent row, we look at the previous row and replace each 
//occurrence of 0 with 01, and each occurrence of 1 with 10. 
//
// 
// For example, for n = 3, the 1ˢᵗ row is 0, the 2ⁿᵈ row is 01, and the 3ʳᵈ row 
//is 0110. 
// 
//
// Given two integer n and k, return the kᵗʰ (1-indexed) symbol in the nᵗʰ row 
//of a table of n rows. 
//
// 
// Example 1: 
//
// 
//Input: n = 1, k = 1
//Output: 0
//Explanation: row 1: 0
// 
//
// Example 2: 
//
// 
//Input: n = 2, k = 1
//Output: 0
//Explanation: 
//row 1: 0
//row 2: 01
// 
//
// Example 3: 
//
// 
//Input: n = 2, k = 2
//Output: 1
//Explanation: 
//row 1: 0
//row 2: 01
// 
//
// 
// Constraints: 
//
// 
// 1 <= n <= 30 
// 1 <= k <= 2n - 1 
// 
//
// Related Topics 位运算 递归 数学 👍 193 👎 0

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
//func kthGrammar(n int, k int) int {
//	return rec([]int{0}, 1, n, k)
//}
//
//func rec(row []int, currentN, targetN, k int) int { 
//	if currentN >= targetN {
//		if len(row) >= k {
//			return row[k-1]
//		} else {
//		return 0
//		}
//	}
//	var next []int
//	for _, num := range row {
//		if num == 0 {
//			next = append(next, []int{0, 1}...)
//		} else {
//			next = append(next, []int{1, 0}...)
//		}
//	}
//	currentN += 1
//	return rec(next, currentN, targetN, k) 
//}
//leetcode submit region end(Prohibit modification and deletion)


func kthGrammar(n int, k int) int {
	fmt.Println(n, k)
	if n == 1 {
		return 0
	}
	return (1 - k % 2) ^ kthGrammar(n-1, (k+1)/2)
} 

func main() {
	fmt.Println(kthGrammar(3, 1))
}