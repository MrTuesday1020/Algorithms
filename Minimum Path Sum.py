"""
https://leetcode.com/problems/minimum-path-sum/
"""

def minPathSum(grid):
	m = len(grid)
	n = len(grid[0])
	
	opt = []
	for i in range(m):
		opt.append([0] * n)
		
	for i in range(m):
		for j in range(n):
			if i == 0:
				opt[i][j] = opt[i][j-1] + grid[i][j]
			elif j == 0:
				opt[i][j] = opt[i-1][j] + grid[i][j]
			else:
				opt[i][j] = min(opt[i-1][j], opt[i][j-1]) + grid[i][j]
				
	return opt[-1][-1]
	
grid = [
	[1,3,1],
	[1,5,1],
	[4,2,1]
]
print(minPathSum(grid))