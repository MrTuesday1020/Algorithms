package main

func setZeroes(matrix [][]int)  {
	rowmap := make(map[int]bool)
	colmap := make(map[int]bool)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				rowmap[i] = true
				colmap[j] = true
			}
		}
	}
	
	for i := range matrix {
		for j := range matrix[i] {
			if rowmap[i] || colmap[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	
}