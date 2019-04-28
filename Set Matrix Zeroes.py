"""
https://leetcode.com/problems/set-matrix-zeroes/solution/
"""

def setZeroes(matrix):
	m = len(matrix)
	n = len(matrix[0])
	
	for i in range(m):
		for j in range(n):
			if matrix[i][j] == 0:
				for p in range(m):
					matrix[p][j] = None if matrix[p][j] != 0 else 0
					
				for q in range(n):
					matrix[i][q] = None if matrix[i][q] != 0 else 0
					
	for i in range(m):
		for j in range(n):
			if matrix[i][j] == None:
				matrix[i][j] = 0
				
	print(matrix)
	
	
matrix = [
	[1,1,1],
	[1,0,1],
	[1,1,1]
]
setZeroes(matrix)