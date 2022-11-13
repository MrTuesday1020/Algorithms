package main

func isValidSudoku(board [][]byte) bool {
	rowmap, colmap, boxmap := initmaps(), initmaps(), initmaps()
	for i, row := range board {
		for j, char := range row {
			if char == '.' {
				continue
			}
			// check row
			if rowmap[i][char] {
				//fmt.Println("row", i , j)
				return false
			} else {
				rowmap[i][char] = true
			}
			// check column
			if colmap[j][char] {
				//fmt.Println("column", i , j)
				return false
			} else {
				colmap[j][char] = true
			}
			// check samll box
			idx := i / 3 * 3 + j / 3
			if boxmap[idx][char] {
				//fmt.Println("box", i , j, idx)
				return false
			} else {
				boxmap[idx][char] = true
			}
		}
	}
	return true
}

func initmaps() []map[byte]bool {
	maps := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		maps[i] = make(map[byte]bool)
	}
	return maps
}

func main() {
	
}