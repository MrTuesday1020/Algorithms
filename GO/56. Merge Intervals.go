package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	
	stack := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		lastRange := stack[len(stack)-1]
		currentRange := intervals[i]
		if lastRange[1] < currentRange[0] {
			// [1,4], [5,6] -> [1,4], [5,6]
			stack = append(stack, currentRange)
			
		} else if lastRange[1] < currentRange[1] {
			// [1,4], [2,5] -> [1,5]
			newRange := []int{lastRange[0], currentRange[1]}
			stack[len(stack)-1] = newRange
		} else {
			// [1,4], [2,3] -> [1,4]
			continue
		}
	}
	return stack
}

func main() {
	fmt.Println(merge([][]int{
		[]int{1,3},
		[]int{2,6},
		[]int{8,10},
		[]int{15,18},
	}))
}