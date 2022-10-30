package main

import "fmt"

var ans [][]int

func permute(nums []int) [][]int {
	ans = make([][]int, 0)
	dfs(nums, []int{})
	return ans
}

func dfs(nums []int, permute []int) {
	if len(nums) == 0 {
		ans = append(ans, permute)
	}
	for i := 0; i < len(nums); i++ {
		newPermute := append(permute, nums[i])
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		newNums = append(newNums[:i], newNums[i+1:]...)
		dfs(newNums, newPermute)
	}
}

func main() {
	fmt.Println(permute([]int{1,2,3}))
}