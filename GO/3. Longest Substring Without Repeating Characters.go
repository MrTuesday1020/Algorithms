package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	stack := []byte{}
	countMap := map[byte]int{}   // key char  value count
	maxLength, currentLength := 0, 0
	
	for i := range s {
		for countMap[s[i]] != 0 {
			// s[i] already exist
			countMap[stack[0]] -= 1		// first element count -1
			stack = stack[1:] 			// pop first element
			currentLength -=1			// current length -1
		}
		
		stack = append(stack, s[i])		// push to stack
		countMap[s[i]] += 1				// s[i] count +1
		currentLength += 1				// current length +1
		if currentLength > maxLength {	// compare to get max length
			maxLength = currentLength
		}
	}
	
	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}