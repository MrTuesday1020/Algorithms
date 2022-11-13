package main

func maxProduct(nums []int) int {
	length := len(nums)
	dpmax := make([]int, length)
	dpmin := make([]int, length)
	dpmax[0], dpmin[0] = nums[0], nums[0]
	ans := nums[0]
	for i := 1; i < length; i++ {
		dpmax[i] = max(nums[i] * dpmax[i-1], max(nums[i], dpmin[i-1] * nums[i]))
		dpmin[i] = min(nums[i] * dpmin[i-1], min(nums[i], dpmax[i-1] * nums[i]))
		ans = max(ans, dpmax[i])
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	
}