package main

func moveZeroes(nums []int)  {
	left, right, n := 0, 0, len(nums)
	//fmt.Println("start",left, right, nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
		//fmt.Println("move ",left, right, nums)
	}
}

func main() {
	
}