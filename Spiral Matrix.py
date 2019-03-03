"""
https://leetcode.com/problems/spiral-matrix/
"""

def spiralOrder(matrix):
	if not matrix:
		return []
		
	ans = []
	while matrix:
		# first row
		if matrix:
			ans += matrix.pop(0)
		# last column
		if matrix:
			for row in matrix:
				ans.append(row.pop())
			matrix = [row for row in matrix if row]
		# last row
		if matrix:
			ans += reversed(matrix.pop())
		# first column
		if matrix:
			for row in reversed(matrix):
				ans.append(row.pop(0))
			matrix = [row for row in matrix if row]
	return ans
	
matrix = [
[7],[9],[6]
]
print(spiralOrder(matrix))