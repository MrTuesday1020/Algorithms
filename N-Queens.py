"""
https://leetcode.com/problems/n-queens/
"""

def solveNQueens(n):
	grids = []
	for i in range(n):
		grids.append(['.']*n)
	res = []
	backtracking(grids, 0, res, n)
	return res
	
def backtracking(grids, row, res, n):
	if row == n:
		solution = []
		for line in grids:
			l = ""
			for item in line:
				l += item
			solution.append(l)
		res.append(solution)
		return
	
	for col in range(n):
		if isValid(grids, row, col, n):
			grids[row][col] = 'Q'
			backtracking(grids, row + 1, res, n)
			grids[row][col] = '.'
			
def isValid(grids, row, col, n):			
	#check if the column had a queen before.
	for i in range(n):
		if grids[i][col] == 'Q' and i != row:
			return False
	#check if the 45° diagonal had a queen before.
	i = row - 1
	j = col - 1
	while i >= 0 and j >= 0:
		if grids[i][j] == 'Q':
			return False
		else:
			i -= 1
			j -= 1
	#check if the 135° diagonal had a queen before.
	i = row - 1
	j = col + 1
	while i >= 0 and j < n:
		if grids[i][j] == 'Q':
			return False
		else:
			i -= 1
			j += 1
			
	return True
	
print(solveNQueens(4))