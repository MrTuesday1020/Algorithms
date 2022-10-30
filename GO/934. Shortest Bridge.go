package main

import "fmt"

var board [][]int
var newboard [][]int
var length int
var done bool

func shortestBridge(grid [][]int) int {
	length = 0
	done = false
	board = make([][]int, 0)
	newboard = make([][]int, 0)
	
	found := false
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 1 {
				findFirstIsland(grid, i, j)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	
	for !done {
		length += 1
		for i := 0; i < len(board); i++ {
			m := board[i][0]
			n := board[i][1]
			// expand in 4 directions
			if expandBoard(grid, m, n-1) ||
				expandBoard(grid, m, n+1) ||
				expandBoard(grid, m+1, n) ||
				expandBoard(grid, m-1, n) {
				done = true 
			} 			
		}
		board = newboard
		newboard = make([][]int, 0)
	} 

	return length-1
}

func findFirstIsland(grid [][]int, i, j int) {
	if i == len(grid) || j == len(grid) || i == -1 || j == -1{
		return
	}
	if grid[i][j] == 1 {
		board = append(board, []int{i, j})
		grid[i][j] = 2
		// move up side
		findFirstIsland(grid, i-1, j) 
		// move down side
		findFirstIsland(grid, i+1, j) 
		// move right side
		findFirstIsland(grid, i, j+1) 
		// move left side
		findFirstIsland(grid, i, j-1) 
	}
}

func expandBoard(grid [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid) {
		return false
	}
	if grid[i][j] == 1 {
		// done
		return true
	}
	if grid[i][j] == 0 {
		// expand new board
		grid[i][j] = 2 + length
		newboard = append(newboard, []int{i, j})
	}
	return false
}

func main() {
	var grid [][]int
	grid = append(grid, []int{0,0,1,0,1})
	grid = append(grid, []int{0,1,1,0,1})
	grid = append(grid, []int{0,1,0,0,1})
	grid = append(grid, []int{0,0,0,0,0})
	grid = append(grid, []int{0,0,0,0,0})
	fmt.Println(shortestBridge(grid))
}