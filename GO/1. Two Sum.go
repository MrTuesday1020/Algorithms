package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // key num   value index
	for i := range nums {
		m[nums[i]] = i
	}
	
	for i := range nums {
		diff := target - nums[i]
		if id, ok := m[diff]; ok && id != i{
			return []int{i, id}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3,2,4}, 6))
}