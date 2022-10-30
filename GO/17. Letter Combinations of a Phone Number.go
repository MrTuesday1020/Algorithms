package main

import "fmt"

var num2char = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

var ret []string

func letterCombinations(digits string) []string {
	ret = []string{}
	if len(digits) == 0 {
		return ret
	}
	backtracking(digits, 0, "")
	return ret
}

func backtracking(digits string, id int, combination string) {
	if id == len(digits) {
		ret = append(ret, combination)
	} else {
		chars := num2char[digits[id]]
		for i := 0; i < len(chars); i++ {
			backtracking(digits, id + 1, combination + string(chars[i]))
		}
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations(""))
}