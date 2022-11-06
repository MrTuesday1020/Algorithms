package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	ans := 0
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		heights := make([]int, cols)
		for j := 0; j < cols; j++ {
			heights[j] = computeHight(matrix, i, j)
		}
		area := maxArea(heights)
		ans = max(ans, area)
	}
	return ans
}

func computeHight(matrix [][]byte, i, j int) int {
	for m := i; m >= 0; m-- {
		if matrix[m][j] == '0' {
			return i - m
		}
	}
	return i + 1
}

func maxArea(heights []int) int {
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
	input := [][]byte{
		[]byte{'1','0','1','0','0'},
		[]byte{'1','0','1','1','1'},
		[]byte{'1','1','1','1','1'},
		[]byte{'1','0','0','1','0'},
	}
	fmt.Println(maximalRectangle(input))
}