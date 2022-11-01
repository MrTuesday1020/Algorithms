package main

import (
	"fmt"
	"sort"
	"strings"
)

func reorganizeString(s string) string {
	// order by count of chars desc
	length := len(s)
	count := make([][]int, 26)
	for i := 0; i < 26; i++ {
		count[i] = []int{i, 0}  // [char, count]
	}
	for i := 0; i < length; i++ {
		id := int(s[i] - 'a')
		count[id][1] += 1
	}
	sort.Slice(count, func(i, j int) bool {
		return count[i][1] > count[j][1]
	})
	max := count[0][1]
	if max > (length + 1) / 2 {
		return ""
	}
	
	// loop char count map desc, append to bucket in order
	bucket := make([]string, max)
	total := 0
	for i := 0; i < 26; i++ {
		char := byte(count[i][0] + 'a')
		cnt := count[i][1]
		if cnt == 0 {
			break
		}
		for j := 0; j < cnt; j++ {
			bucket[total%max] += string(char)
			total += 1
		}
	}
	
	return strings.Join(bucket, "")
}

func main() {
	fmt.Println(reorganizeString("aabbbc"))
}