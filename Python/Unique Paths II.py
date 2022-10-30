"""
https://leetcode.com/problems/unique-paths-ii/
"""

def uniquePathsWithObstacles(obstacleGrid):
	m = len(obstacleGrid)
	n = len(obstacleGrid[0])
	opt = []
	for i in range(m):
		opt.append([0]*n)
	
	for i in range(m):
		for j in range(n):
			# start
			if i == 0 and j == 0:
				if obstacleGrid[i][j] == 1:
					return 0
				opt[i][j] = 1
			# obstacle
			elif obstacleGrid[i][j] == 1:
				opt[i][j] = 0
			# first row
			elif i == 0:
				opt[i][j] = opt[i][j-1]
			# first column
			elif j == 0:
				opt[i][j] = opt[i-1][j]
			# rest
			else:
				opt[i][j] = opt[i-1][j] + opt[i][j-1]
					
	return opt[-1][-1]
	
	
obstacleGrid = [
	[0,0,0],
	[0,1,0],
	[0,0,0]
]
print(uniquePathsWithObstacles(obstacleGrid))