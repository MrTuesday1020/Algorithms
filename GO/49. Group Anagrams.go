package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		sortedStr := sortStr(strs[i])
		if _, ok := m[sortedStr]; ok {
			m[sortedStr] = append(m[sortedStr], strs[i])
		} else {
			m[sortedStr] = []string{strs[i]}
		}
	}

	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)	
	}
	return ans
}

func sortStr(in string) string {
	var arr []byte
	for i := 0; i < len(in); i++ {
		arr = append(arr, in[i])
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return string(arr)
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))	
}