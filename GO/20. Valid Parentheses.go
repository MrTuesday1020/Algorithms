package main

import "fmt"

func isValid(s string) bool {
	var stack []rune
	for _, char := range s {
		if len(stack) == 0 || char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if char == ')' && stack[len(stack)-1] == '(' ||
				char == '}' && stack[len(stack)-1] == '{' ||
				char == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, char)
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}