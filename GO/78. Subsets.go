package main

import "fmt"

func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, set)
	}
	return
}

func main() {
	fmt.Println(4>>2&1) // 100 -> 1 -> 1 & 1 -> 1
	fmt.Println(4>>1&1) // 100 -> 10 -> 10 & 01 -> 0
	fmt.Println(4>>0&1) // 100 -> 100 > 100 & 001 -> 0
	
	fmt.Println(5>>2&1) // 101 -> 1 -> 1 & 1 -> 1
	fmt.Println(5>>1&1) // 101 -> 10 -> 10 & 01 -> 0
	fmt.Println(5>>0&1) // 101 -> 101 > 101 & 001 -> 1
}