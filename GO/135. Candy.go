package main

import "fmt"

func candy(ratings []int) int {
	l := len(ratings)
	if l <= 1 {
		return l
	}
	
	candies1 := make([]int, l)
	candies2 := make([]int, l)
	
	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			candies1[i] = candies1[i-1] + 1
		}
	}
	
	for i := l-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies2[i] = candies2[i+1] + 1
		}
	}
	
	ans := 0
	for i := 0; i < l; i++ {
		ans += max(candies1[i], candies2[i]) + 1
	}
	return ans
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	fmt.Println(candy([]int{1,2,2}))
	fmt.Println(candy([]int{1,3,2,2,1}))
	fmt.Println(candy([]int{1,2,87,87,87,2,1}))
}