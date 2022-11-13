package main

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				count += 1
				dfs(grid, i, j, visited)
			}
		}
	}
	
	return count
}

func dfs(grid [][]byte, i, j int, visited [][]bool) {
	visited[i][j] = true
	// up
	if i > 0 && grid[i-1][j] == '1' && !visited[i-1][j] {
		dfs(grid, i-1, j, visited)
	}
	// down
	if i < len(grid)-1 && grid[i+1][j] == '1' && !visited[i+1][j] {
		dfs(grid, i+1, j, visited)
	}
	// left
	if j > 0 && grid[i][j-1] == '1' && !visited[i][j-1] {
		dfs(grid, i, j-1, visited)
	}
	// right
	if j < len(grid[0])-1 && grid[i][j+1] == '1' && !visited[i][j+1] {
		dfs(grid, i, j+1, visited)
	}
}

func main() {
	
}