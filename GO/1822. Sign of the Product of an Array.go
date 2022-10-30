package main

func arraySign(nums []int) int {
	ans := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			ans *= -1
		}
	}
	return ans
}

func main() {
	
}