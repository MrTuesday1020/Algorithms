package main

func subsetsWithDup(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return i < j
	})
	n := len(nums)
	
	var ans [][]int
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		flag := true
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					flag = false
					break
				}
				set = append(set, v)
			}
		}
		if flag {
			ans = append(ans, set)
		}
	}
	
	return ans
}

func main() {
	
}