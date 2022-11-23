package main

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	var sCount, pCount [26]int
	for i := range p {
		pCount[p[i]-'a'] += 1
		sCount[s[i]-'a'] += 1
	}
	ans := []int{}
	if pCount == sCount {
		ans = append(ans, 0)
	}
	for i := len(p); i < len(s); i++ {
		// move window forward
		sCount[s[i-len(p)]-'a'] -= 1 // remove previous element
		sCount[s[i]-'a'] += 1 // add next element
		if pCount == sCount {
			ans = append(ans, i-len(p)+1) // add current first element index
		}
	}
	return ans
}

func main() {
	
}