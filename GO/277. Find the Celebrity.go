package main

import "fmt"

var people [][]int

/**
* The knows API is already defined for you.
*     knows := func(a int, b int) bool
*/

func knows(a, b int) bool {
	if len(people) > a && len(people[a]) > b {
		return people[a][b] == 1
	}
	return false
}

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		ans := 0
		for i := 0; i < n; i++ {
			if ans == i {
				continue
			}
			
			if knows(ans, i) || !knows(i, ans) {
				// ans knows i or i does not know ans
				ans = i
			}
		}
		
		fmt.Println(ans)
		
		for i := 0; i < n; i++ {
			if ans == i {
				continue
			}
			if knows(ans, i) || !knows(i, ans) {
				// ans knows i or i does not know ans
				return -1
			}
		}
		
		return ans
	}
}

func main() {
	people = [][]int{
		[]int{1,1,0},
		[]int{0,1,0},
		[]int{1,1,1},
	}
	fn := solution(knows)
	fmt.Println(fn(3))
	
	people = [][]int{
		[]int{1,0,1},
		[]int{1,1,0},
		[]int{0,1,1},
	}
	fn = solution(knows)
	fmt.Println(fn(3))
}