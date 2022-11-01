package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited := make([][]bool, m)
			for k := range visited {
				visited[k] = make([]bool, n)
			}
			if dfs(board, visited, word, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visited [][]bool, str string, i, j int) bool {
	if str == "" { // done
		return true
	}
	if i > len(board)-1 || j > len(board[0])-1 || i < 0 || j < 0 { // out of board
		return false
	}
	if board[i][j] != str[0] || visited[i][j] { // board[i][j] not match or visited
		return false
	}
	visited[i][j] = true
	defer func() { visited[i][j] = false }() // IMPORT!!!!!
	return dfs(board, visited, str[1:], i-1, j) || 
	dfs(board, visited, str[1:], i, j-1) ||
	dfs(board, visited, str[1:], i+1, j) ||
	dfs(board, visited, str[1:], i, j+1)
}

func main() {
	board := [][]byte{
		[]byte{'A','B','C','E'},
		[]byte{'S','F','E','S'},
		[]byte{'A','D','E','E'},
	}
	fmt.Println(exist(board, "ABCESEEEFS"))
}