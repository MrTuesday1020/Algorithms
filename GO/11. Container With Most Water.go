package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		l := right - left
		h := min(height[left], height[right])
		maxArea = max(h * l, maxArea)
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return maxArea
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}