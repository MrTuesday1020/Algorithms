package main

import "fmt"

func minWindow(s string, t string) string {
	sourceM := make(map[byte]int)
	targetM := make(map[byte]int)
	for i := range t {
		targetM[t[i]] += 1
	}
	length := len(s)
	
	got := false
	l, r, begin, end := 0, 0, 0, 0
	for l <= r && r < length + 1 {
		if checkContain(sourceM, targetM) {
			begin, end, got = checkAns(l, r, begin, end, got)
			if l < length {
				sourceM[s[l]] -= 1
			}
			l += 1
		} else {
			if r < length {
				sourceM[s[r]] += 1
			}
			r += 1
		}
	}
	
	return s[begin:end]
}

func checkAns(l, r, begin, end int, got bool) (int, int, bool) {
	if !got {
		return l, r, true
	} 
	if r - l < end - begin {
		return l, r, true
	}
	return begin, end, true
}

func checkContain(sourceM, targetM map[byte]int) bool {	
	for k, v := range targetM {
		if sourceM[k] < v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}