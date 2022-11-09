package main

import "fmt"

func convert(s string, numRows int) string {
	strs := make([]string, numRows)
	for i := 0; i < len(s); i++ {
		idx := i % (numRows + numRows - 2)
		if idx < numRows {
			strs[idx] += string(s[i])
		} else {
			idx = numRows + numRows - 2 -idx
			strs[idx] += string(s[i])
		}
	}
	
	//fmt.Println(strs)
	
	var ans string
	for i := 0; i < numRows; i++ {
		ans += strs[i]
	}
	return ans
}

func main() {
	fmt.Println(convert("leetcode", 4))
}