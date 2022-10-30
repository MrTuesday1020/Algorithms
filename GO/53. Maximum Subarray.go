package main

import "fmt"

func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		ans = max(ans, nums[i])
	}
	return ans
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}