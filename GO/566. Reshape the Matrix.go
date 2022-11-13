package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m * n != r * c {
		return mat
	}
	
	var ans [][]int
	for i := 0; i < r; i++ {
		ans = append(ans, make([]int, c))
	}
	
	idx := 0
	for _, row := range mat {
		for _, num := range row {
			i := idx / c
			j := idx % c
			ans[i][j] = num
			idx += 1
		}
	}
	
	return ans
}

func main() {
	
}