def exist(board, word):
	m = len(board)
	n = len(board[0])
	
	for i in range(m):
		for j in range(n):
			if search(board, i, j, word, m, n):
				return True
	return False
				
def search(board, i, j, word, m, n):
	if len(word) == 0:
		return True
		
	if i < 0 or i >= m or j < 0 or j >= n or word[0]!=board[i][j]:
		return False
		
	tmp = board[i][j]  # first character is found, check the remaining part
	board[i][j] = None  # avoid visit agian 
	ans = search(board, i-1, j, word[1:], m, n) or search(board, i, j-1, word[1:], m, n) or search(board, i+1, j, word[1:], m, n) or search(board, i, j+1, word[1:], m, n)
	board[i][j] = tmp		
	return ans
	
board = [["C","A","A"],["A","A","A"],["B","C","D"]]
word = "AAB"

print(exist(board, word))