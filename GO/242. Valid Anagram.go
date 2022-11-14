package main

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, char := range s {
		m[char] += 1
	}
	for _, char := range t {
		m[char] -= 1
		if m[char] < 0 {
			return false
		}
	}
	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	
}