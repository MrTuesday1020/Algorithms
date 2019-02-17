"""
https://leetcode.com/problems/valid-sudoku/
"""

def isValidSudoku(board: 'List[List[str]]') -> 'bool':
	rows = [{},{},{},{},{},{},{},{},{}]
	columns = [{},{},{},{},{},{},{},{},{}]
	boxes = [{},{},{},{},{},{},{},{},{}]
	
	for i in range(9):
		for j in range(9):
			if board[i][j] != ".":
				# check row
				if board[i][j] in rows[i]:
					return False
				else:
					rows[i].update({board[i][j]: 1})
					
				# check column
				if board[i][j] in columns[j]:
					return False
				else:
					columns[j].update({board[i][j]: 1})
					
				# check box
				boxid = i // 3 * 3 + j // 3 
				if board[i][j] in boxes[boxid]:
					return False
				else:
					boxes[boxid].update({board[i][j]: 1})
					
	return True
	
	
board = [
	["8","3",".",".","7",".",".",".","."],
	["6",".",".","1","9","5",".",".","."],
	[".","9","8",".",".",".",".","6","."],
	["8",".",".",".","6",".",".",".","3"],
	["4",".",".","8",".","3",".",".","1"],
	["7",".",".",".","2",".",".",".","6"],
	[".","6",".",".",".",".","2","8","."],
	[".",".",".","4","1","9",".",".","5"],
	[".",".",".",".","8",".",".","7","9"]
]

print(isValidSudoku(board))