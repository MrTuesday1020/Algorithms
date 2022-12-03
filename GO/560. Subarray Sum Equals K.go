package main

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	pre, count := 0, 0
	for i := range nums {
		pre += nums[i]
		if v, ok := m[pre - k]; ok {
			count += v
		}
		m[pre] += 1
	}
	return count
}

func main() {
	
}