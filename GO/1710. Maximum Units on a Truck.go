package main

import (
	"sort"
	"fmt"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
		
	ans := 0
	i := 0
	for truckSize > 0 && i < len(boxTypes) {
		if truckSize > boxTypes[i][0] {
			ans += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			ans += truckSize * boxTypes[i][1]
			truckSize = 0
		}
		i += 1
	}
	return ans
}

func main() {
	boxes := [][]int{
		[]int{5,10},
		[]int{2,5},
		[]int{4,7},
		[]int{3,9},
	}
	fmt.Println(maximumUnits(boxes, 10))
	
	
	boxes = [][]int{
		[]int{1,3},
		[]int{2,2},
		[]int{3,1},
	}
	fmt.Println(maximumUnits(boxes, 4))
}