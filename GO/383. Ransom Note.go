package main

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, char := range magazine {
		m[char] += 1
	}
	for _, char := range ransomNote {
		m[char] -= 1
		if m[char] < 0 {
			return false
		}
	}
	return true
}

func main() {
	
}