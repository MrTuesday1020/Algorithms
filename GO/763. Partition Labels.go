package main

func partitionLabels(s string) []int {
	lastPos := make(map[rune]int)
	for i, char := range s {
		lastPos[char] = i
	}
	
	var ans []int
	start, end := 0, 0
	for i, char := range s {
		if lastPos[char] > end {
			end = lastPos[char]
		}
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}
	return ans
}

func main() {
	
}