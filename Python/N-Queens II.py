"""
https://leetcode.com/problems/n-queens-ii/
"""

"""
 * don't need to actually place the queen,
 * instead, for each row, try to place without violation on
 * col/ diagonal1/ diagnol2.
 * trick: to detect whether 2 positions sit on the same diagnol:
 * if delta(col, row) equals, same diagnol1;
 * if sum(col, row) equals, same diagnal2.
"""

def totalNQueens(n):
	cols = [False]*n
	diag1 = [False]*2*n
	diag2 = [False]*2*n
	count = backtracking(cols, diag1, diag2, 0, n, 0)
	return count
	
	
def backtracking(cols, diag1, diag2, row, n, count):
	for col in range(n):
		d1 = row - col # 利用python的list的特性，d1<0时也有用
		d2 = row + col
		if cols[col] or diag1[d1] or diag2[d2]:
			continue
		if row == n-1:
			count += 1
		else:
			cols[col] = diag1[d1] = diag2[d2] = True
			count = backtracking(cols, diag1, diag2, row+1, n, count)
			cols[col] = diag1[d1] = diag2[d2] = False
			
	return count
			
print(totalNQueens(4))