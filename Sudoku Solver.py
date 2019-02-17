"""
https://leetcode.com/problems/sudoku-solver/
https://www.youtube.com/watch?v=JzONv5kaPJM
"""

def solveSudoku(board: 'List[List[str]]') -> 'None':
	if not board:
		return
		
	# build dicts
	rows = [set(),set(),set(),set(),set(),set(),set(),set(),set()]
	columns = [set(),set(),set(),set(),set(),set(),set(),set(),set()]
	boxes = [set(),set(),set(),set(),set(),set(),set(),set(),set()]
		
	for i in range(9):
		for j in range(9):
			if board[i][j] != ".":
				rows[i].add(board[i][j])
				columns[j].add(board[i][j])
				boxid = i // 3 * 3 + j // 3 
				boxes[boxid].add(board[i][j])
	
	digits = ["1","2","3","4","5","6","7","8","9"]
	def helper(board,rows,columns,boxes):
		for i in range(9):
			for j in range(9):
				if board[i][j] == ".":
					for digit in digits:
						# check valid
						boxid = i // 3 * 3 + j // 3 
						if digit not in rows[i] and digit not in columns[j] and digit not in boxes[boxid]:
							board[i][j] = digit
							# update sets
							rows[i].add(digit)
							columns[j].add(digit)
							boxes[boxid].add(digit)
							
							if helper(board,rows,columns,boxes):
								return True
							else:
								board[i][j] = "."
								# update sets
								rows[i].remove(digit)
								columns[j].remove(digit)
								boxes[boxid].remove(digit)
					return False
		return True
		
	helper(board,rows,columns,boxes)
	print(board)
	
board = [
	["5","3",".",".","7",".",".",".","."],
	["6",".",".","1","9","5",".",".","."],
	[".","9","8",".",".",".",".","6","."],
	["8",".",".",".","6",".",".",".","3"],
	["4",".",".","8",".","3",".",".","1"],
	["7",".",".",".","2",".",".",".","6"],
	[".","6",".",".",".",".","2","8","."],
	[".",".",".","4","1","9",".",".","5"],
	[".",".",".",".","8",".",".","7","9"]
]
solveSudoku(board)