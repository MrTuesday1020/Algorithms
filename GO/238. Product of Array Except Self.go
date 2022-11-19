package main

func productExceptSelf(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	suffix := make([]int, length)
	ans[0] = 1
	for i := 1; i < length; i++ {
		ans[i] = nums[i-1] * ans[i-1] // 前缀之积
	}
	suffix[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]  // 后缀之积
	}
	for i := 0; i < length; i++ {
		ans[i] = ans[i] * suffix[i]  // 最后结果
	}
	return ans
}

func main() {
	
}