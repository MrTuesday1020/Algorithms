package main

import "fmt"

var ret []string

func generateParenthesis(n int) []string {
	ret = []string{}
	recursion(n, 0 , 0 , "")
	return ret
}

func recursion(n, left, right int, substr string) {
	if n == left && n == right {
		ret = append(ret, substr)
		return
	}
	if left == right {
		recursion(n, left + 1, right, substr + "(")
	} else {
		if left < n {
			recursion(n, left + 1, right, substr + "(")
		}
		recursion(n, left, right + 1, substr + ")")
	}
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(3))
}