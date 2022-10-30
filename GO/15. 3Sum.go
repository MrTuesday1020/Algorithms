package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	retMap := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		sum := 0 - nums[i]
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left] + nums[right] == sum {
				// find a solution
				key := fmt.Sprint("%d_%d_%d", nums[i], nums[left], nums[right])
				if !retMap[key] { // remove dup
					ret = append(ret, []int{nums[i], nums[left], nums[right]})
					retMap[key] = true
				}
				left += 1
			} else if nums[left] + nums[right] < sum {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}