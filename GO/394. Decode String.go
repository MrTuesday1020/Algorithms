package main

//Given an encoded string, return its decoded string. 
//
// The encoding rule is: k[encoded_string], where the encoded_string inside the 
//square brackets is being repeated exactly k times. Note that k is guaranteed to 
//be a positive integer. 
//
// You may assume that the input string is always valid; there are no extra 
//white spaces, square brackets are well-formed, etc. Furthermore, you may assume 
//that the original data does not contain any digits and that digits are only for 
//those repeat numbers, k. For example, there will not be input like 3a or 2[4]. 
//
// The test cases are generated so that the length of the output will never 
//exceed 10⁵. 
//
// 
// Example 1: 
//
// 
//Input: s = "3[a]2[bc]"
//Output: "aaabcbc"
// 
//
// Example 2: 
//
// 
//Input: s = "3[a2[c]]"
//Output: "accaccacc"
// 
//
// Example 3: 
//
// 
//Input: s = "2[abc]3[cd]ef"
//Output: "abcabccdcdcdef"
// 
//
// 
// Constraints: 
//
// 
// 1 <= s.length <= 30 
// s consists of lowercase English letters, digits, and square brackets '[]'. 
// s is guaranteed to be a valid input. 
// All the integers in s are in the range [1, 300]. 
// 
//
// Related Topics 栈 递归 字符串 👍 1331 👎 0

import (
	"fmt"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
var i int
var source string

func decodeString(s string) string {
	i = 0
	source = s
	return getString()
}

func getString() string {
	if i == len(source) {
		return ""
	}
	if source[i] == ']'  {
		// jump ]
		i += 1
		return ""
	}
	
	num := 1
	pre := ""
	if source[i] >= '0' && source[i] <= '9' {
		num = getNumber()
		next := getString()
		pre = strings.Repeat(next, num)
	} else if source[i] >= 'a' && source[i] <= 'z' || source[i] >= 'A' && source[i] <= 'Z' {
		pre = string(source[i])
		i += 1
	}
	return pre + getString()
}

func getNumber() int {
	str := ""
	for source[i] >= '0' && source[i] <= '9' {
		str += string(source[i])
		i += 1
	}
	i += 1 // jump [
	num, _ := strconv.Atoi(str)
	return num
}

//leetcode submit region end(Prohibit modification and deletion)


func main() {
	fmt.Println(decodeString("100[leetcode]"))
}