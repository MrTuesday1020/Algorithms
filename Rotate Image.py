"""
https://leetcode.com/problems/rotate-image/
"""

def rotate(matrix):
	"""
	Do not return anything, modify matrix in-place instead.
	"""
	
	#matrix[::] = zip(*matrix[::-1])
	
	n = len(matrix)
	
	for i in range(n//2):
		for j in range(n-n//2):
			matrix[i][j], matrix[~j][i], matrix[~i][~j], matrix[j][~i] = \
			 matrix[~j][i], matrix[~i][~j], matrix[j][~i], matrix[i][j]
			
	print(matrix)
	
matrix = [
	[ 5, 1, 9,11],
	[ 2, 4, 8,10],
	[13, 3, 6, 7],
	[15,14,12,16]
]
rotate(matrix)