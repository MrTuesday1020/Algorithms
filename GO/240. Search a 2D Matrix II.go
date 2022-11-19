package main

func searchMatrix(matrix [][]int, target int) bool {
	col := len(matrix)-1
	row := 0
	for col >=0 && row < len(matrix[0]) {
		if matrix[col][row] == target {
			return true
		} else if matrix[col][row] > target {
			col -= 1
		} else {
			row += 1
		}
	}
	return false
}

func main() {
	
}