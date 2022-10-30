package main

import "fmt"

func trap(height []int) int {
	length := len(height)
	if length < 3 {
		return 0
	}

	leftMax := make([]int, length)
	rightMax := make([]int, length)
	leftMax[0] = height[0]
	rightMax[length-1] = height[length-1]
	ans := 0
	
	for i := 1; i < length; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := length-2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i := 0; i < length; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}