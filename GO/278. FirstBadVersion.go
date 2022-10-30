//You are a product manager and currently leading a team to develop a new 
//product. Unfortunately, the latest version of your product fails the quality check. 
//Since each version is developed based on the previous version, all the versions 
//after a bad version are also bad. 
//
// Suppose you have n versions [1, 2, ..., n] and you want to find out the 
//first bad one, which causes all the following ones to be bad. 
//
// You are given an API bool isBadVersion(version) which returns whether 
//version is bad. Implement a function to find the first bad version. You should 
//minimize the number of calls to the API. 
//
// 
// Example 1: 
//
// 
//Input: n = 5, bad = 4
//Output: 4
//Explanation:
//call isBadVersion(3) -> false
//call isBadVersion(5) -> true
//call isBadVersion(4) -> true
//Then 4 is the first bad version.
// 
//
// Example 2: 
//
// 
//Input: n = 1, bad = 1
//Output: 1
// 
//
// 
// Constraints: 
//
// 
// 1 <= bad <= n <= 2³¹ - 1 
// 
//
// Related Topics 二分查找 交互 👍 811 👎 0


package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */



var BadVersion int

func firstBadVersion(n int) int {
	return findFirstBadVersion(1,n)
}

func findFirstBadVersion(start, end int) int {
	if start > end {
		return start
	}
	mid := start + ( end - start ) / 2
	if isBadVersion(mid) {
		return findFirstBadVersion(start, mid - 1)
	} 
	return findFirstBadVersion(mid + 1, end)
}

//leetcode submit region end(Prohibit modification and deletion)


func isBadVersion(version int) bool {
	if version >= BadVersion {
		return true
	}
	return false
}


func main() {
	BadVersion = 10
	fmt.Println(firstBadVersion(11))
}