package main

import (
	"fmt"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	horizontalCuts = append([]int{0}, horizontalCuts...)
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append([]int{0}, verticalCuts...)
	verticalCuts = append(verticalCuts, w)
	
	sort.Slice(horizontalCuts, func(i, j int) bool {
		return horizontalCuts[i] < horizontalCuts[j]
	})
	sort.Slice(verticalCuts, func(i, j int) bool {
		return verticalCuts[i] < verticalCuts[j]
	})
	
	maxHight, maxWidth := 0, 0
	for i := 1; i < len(horizontalCuts); i++ {
		height := horizontalCuts[i] - horizontalCuts[i-1]
		maxHight = max(height, maxHight)
	}
	for i := 1; i < len(verticalCuts); i++ {
		width := verticalCuts[i] - verticalCuts[i-1]
		maxWidth = max(width, maxWidth)
	}
	fmt.Println(maxHight, maxWidth)
	
	return (maxHight * maxWidth) % 1000000007
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(maxArea(5, 4, []int{1,2,4}, []int{1,3}))
}