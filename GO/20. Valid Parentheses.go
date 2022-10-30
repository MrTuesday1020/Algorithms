package main

import "fmt"

func isValid(s string) bool {
	var stack []string
	for i := range s {
		current := string(s[i])
		if len(stack) > 0 {
			last := stack[len(stack)-1]
			if (last == "(" && current == ")") || 
			(last == "[" && current == "]") || 
			(last == "{" && current == "}") {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, current)
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}