"""
https://leetcode.com/problems/edit-distance/
https://www.youtube.com/watch?v=MiqoA-yF-0M
"""

def minDistance(word1, word2):
	m, n, opt = len(word1) + 1, len(word2) + 1, []

	for i in range(m):
		opt.append([None]*n)
		
	for i in range(m):
		for j in range(n):
			if i == 0 and j == 0:
				opt[i][j] = 0
			elif i == 0 and j != 0:
				opt[i][j] = j
			elif i != 0 and j == 0:
				opt[i][j] = i
			else:
				if word1[i-1] != word2[j-1]:
					opt[i][j] = min(opt[i-1][j], opt[i][j-1], opt[i-1][j-1]) + 1
				else:
					opt[i][j] = min(opt[i-1][j] + 1, opt[i][j-1] + 1, opt[i-1][j-1])
	return opt[-1][-1]
	
word1 = "zoologicoarchaeologist"
word2 = "zoogeologist"

print(minDistance(word1, word2))