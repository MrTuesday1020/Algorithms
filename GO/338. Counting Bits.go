package main

import "fmt"

func countBits(n int) []int {
	ans := make([]int, n+1)
	ans[0] = 0
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			// 如果正整数 y 是 2 的整数次幂，则 y 的二进制表示中只有最高位是 1，其余都是 0，因此 y & (y−1) = 0
			// 比如 2, 4, 8, 16
			highBit = i
		}
		// i 和 i-highBit 相比 只是多了最高位的一个 1
		ans[i] = ans[i-highBit] + 1
	}
	return ans
}

func main() {
	countBits(10)
}