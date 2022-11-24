package main

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]bool)
	for i := range nums {
		m[nums[i]] = true
	}
	ans := []int{}
	for i := 1; i <= len(nums); i++ {
		if !m[i] {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	
}