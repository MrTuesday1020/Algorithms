package main

func reconstructQueue(people [][]int) [][]int {
	// [[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	// [[7,0],[7,1],[6,1],[5,0],[5,2],[4,4]]
	ans := [][]int{}
	for _, person := range people {
		idx := person[1]
		// insert at position idx
		temp := append([][]int{}, ans[:idx]...)
		temp = append(temp, person)
		temp = append(temp, ans[idx:]...)
		ans = temp
	}
	// [[7,0]]
	// [[7,0],[7,1]]
	// [[7,0],[6,1],[7,1]]
	// [[5,0],[7,0],[6,1],[7,1]]
	// [[5,0],[7,0],[5,2],[6,1],[7,1]]
	// [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
	return ans
}

func main() {
	
}