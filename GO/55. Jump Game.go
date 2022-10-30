package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	
	length := len(nums)
	maxDis := 0
	
	for i := 0; i < length; i++ {
		if i <= maxDis { 						// i is reachable
			maxDis = max(maxDis, i + nums[i]) 	// max distance can be reached
			if maxDis >= length - 1 { 			// already reach the last point
				return true
			}
		} else {
			return false						// i is not reachable, break
		}
	} 
	return false
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(canJump([]int{1,2}))
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
}