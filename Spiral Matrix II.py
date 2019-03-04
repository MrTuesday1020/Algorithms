"""
https://leetcode.com/problems/spiral-matrix-ii/
"""

def generateMatrix(n):
	if n == 0:
		return
	if n == 1:
		return [[1]]
		
	ans = []
	for i in range(n):
		ans.append([0] * n)
	
	layer = 0
	current = 1
	while layer < n / 2:
		# top
		nb = n - layer * 2
		for i in range(nb):
			ans[layer][layer+i] = current
			if current == n * n:
				break
			current += 1
		# right
		nb = n - layer * 2 - 2
		for i in range(nb):
			ans[layer+1+i][n-1-layer] = current
			if current == n * n:
				break
			current += 1
		# bottom
		nb = n - layer * 2
		for i in range(nb):
			ans[n-1-layer][n-layer-i-1] = current
			if current == n * n:
				break
			current += 1
		# left
		nb = n - layer * 2 - 2
		for i in range(nb):
			ans[n-layer-i-2][layer] = current
			if current == n * n:
				break
			current += 1
		layer += 1
		
	return ans
	
print(generateMatrix(4))