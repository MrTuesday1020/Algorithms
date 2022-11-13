package main

func generate(numRows int) [][]int {
	ans := [][]int{
		[]int{1},
	}
	for i := 1; i < numRows; i++ {
		preLevel := ans[len(ans)-1]
		level := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				level[j] = 1
			} else {
				level[j] = preLevel[j-1] + preLevel[j]
			}
		}
		ans = append(ans, level)
	}
	return ans
}

func main() {
	
}