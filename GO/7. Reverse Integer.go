package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ans := 0
	for x != 0 {
		if ans < math.MinInt32 / 10 || ans > math.MaxInt32 / 10 {
			return 0
		}
		digit := x % 10
		x = x / 10
		ans = ans * 10 + digit
	}
	return ans
}

func main() {
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1234567899))
}