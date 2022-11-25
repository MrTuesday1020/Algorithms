package main

func hammingDistance(x int, y int) int {
	s := x ^ y
	ans := 0
	for s > 0 {
		if s & 1 == 1 {
			ans += 1
		}
		s >>= 1
	}
	return ans
}

func main() {
	
}