package main

import (
	"fmt"
	"sort"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		m := fmt.Sprint(nums[i])
		n := fmt.Sprint(nums[j])
		return m+n > n+m
	})
	
	if nums[0] == 0 {
		return "0"
	}
	
	ans := ""
	for i := range nums {
		ans += fmt.Sprint(nums[i])
	}
	return ans
}

func main() {
	fmt.Println(largestNumber([]int{3,30,34,5,9}))
}