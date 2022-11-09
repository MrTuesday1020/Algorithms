package main

func isPalindrome(x int) bool {
	if x < 0 {
	return false
	}
	
	l := 0
	divider := 1
	for x / divider > 0 {
	divider *= 10
	l += 1
	}
	
	y := 0
	for i := 0; i < l/2; i++ {
	y = y * 10 + x % 10
	x /= 10
	}
	
	if l % 2 != 0 {
	x /= 10
	}
	
	return x == y
}

func main() {
	
}