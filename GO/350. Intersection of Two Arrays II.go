package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)    // key num  value count
	for _, num := range nums1 {
		m[num] += 1
	}
	
	var ans []int
	for _, num := range nums2 {
		count := m[num]
		if count >0 {
			ans = append(ans, num)
			m[num] -= 1
		}
	}
	
	return ans
}

func main() {
	
}