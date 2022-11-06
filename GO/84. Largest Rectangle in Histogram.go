package main

import "fmt"

func largestRectangleArea(heights []int) int {
	length := len(heights)
	stack := []int{-1, 0}
	ans := 0
	for i := 1; i < length; i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			width := i - stack[len(stack)-2] - 1
			stack = stack[:len(stack)-1]
			ans = max(ans, height * width)
		}
		stack = append(stack, i)
	}
		
	for i := len(stack)-1; i > 0; i-- {
		height := heights[stack[i]]
		width := length - stack[i-1] - 1
		ans = max(ans, height * width)
	}
	
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(largestRectangleArea([]int{5,4,1,2}))
}