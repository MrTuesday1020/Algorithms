package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = false
	}
	ans := 1
	for _, num := range nums {
		if !m[num] {		// make sure every number is visited once [0(N)]
			length := 1
			m[num] = true 	// num visited
			
			// try to find next num
			next := num + 1
			for {
				if visited, got := m[next]; !visited && got {
					m[next] = true 	// next num visited
					next += 1
					length += 1
				} else {
					break
				}
			}
			
			// try to find pre num
			pre := num - 1
			for {
				if visited, got := m[pre]; !visited && got {
					m[pre] = true 	// pre num visited
					pre -= 1
					length += 1
				} else {
					break
				}
			}

			ans = max(ans, length)
		}
	}
	return ans
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	fmt.Println(longestConsecutive([]int{100,4,200,1,3,2}))
}