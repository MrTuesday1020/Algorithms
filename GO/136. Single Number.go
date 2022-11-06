package main

func singleNumber(nums []int) int {
	num := 0
	for i := range nums {
		num ^= nums[i]
	}
	return num
}

func main() {
	
}