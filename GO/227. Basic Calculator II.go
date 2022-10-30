package main

import (
	"fmt"
	"strconv"
)

func calculate(s string) int {
	var strs []string
	var currentNum string
	for i := range s {
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			strs = append(strs, currentNum)
			strs = append(strs, string(s[i]))
			currentNum = ""
			continue
		}
		if s[i] != ' ' {
			currentNum += string(s[i])
		}
	}
	strs = append(strs, currentNum)
	var nums []int
	var op string
	for i := 0; i < len(strs); i++ {
		// 如果是符号 就给op赋值
		if strs[i] == "+" || strs[i] == "-" || strs[i] == "*" || strs[i] == "/" {
			op = strs[i]
			continue
		}
		// 如果是数字 就计算
		num, _ := strconv.Atoi(strs[i])
		if op == "+" || i == 0{
			nums = append(nums, num)
		} else if op == "-" {
			nums = append(nums, 0-num)
		} else if op == "*" {
			nums[len(nums)-1] *= num 
		} else if op == "/" {
			nums[len(nums)-1] /= num
		} 
	}
	var ret int
	for i := range nums {
		ret += nums[i]
	}
	return ret
}

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
}