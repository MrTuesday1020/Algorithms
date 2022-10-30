package main

import "fmt"

func subArrayRanges(nums []int) int64 {
	if len(nums) == 1 {
		return 0
	}
	var sum int64
	for i := 0; i < len(nums) - 1; i++ {
		min := nums[i]
		max := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
			}
			if nums[j] > max {
				max = nums[j]
			}
			sum += int64(max -min)
		}
	}
	return sum
}

func main() {
	fmt.Print(subArrayRanges([]int{4,-2,-3,4,1}))
}