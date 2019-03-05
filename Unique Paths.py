"""
https://leetcode.com/problems/unique-paths/
"""

def uniquePaths(m, n):
	opt = []
	for i in range(m):
		opt.append([0] * n)
		
	for i in range(m):
		for j in range(n):
			if i == 0 or j == 0:
				opt[i][j] = 1
			else:
				opt[i][j] = opt[i-1][j] + opt[i][j-1]
				
	return opt[-1][-1]
	
print(uniquePaths(3, 2))