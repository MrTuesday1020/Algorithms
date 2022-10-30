"""
https://leetcode.com/problems/search-a-2d-matrix/
"""

def searchMatrix(matrix, target):
	m = len(matrix)
	if m == 0:
		return False
	n = len(matrix[0])
	if n == 0:
		return False
	
	for i in range(m):
		if matrix[i][0] <= target and matrix[i][-1] >= target:
			start, end = 0, n-1
			while start <= end:
				mid = (start + end) // 2
				if matrix[i][mid] == target:
					return True
				elif matrix[i][mid] < target:
					start = mid + 1
				else:
					end = mid - 1
			return False
	
	return False
	
matrix = [
[1]
]
target = 1
print(searchMatrix(matrix, target))