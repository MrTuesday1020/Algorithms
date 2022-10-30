"""
https://leetcode.com/problems/regular-expression-matching/
"""

def isMatch(s, p):
	"""
	:type s: str
	:type p: str
	:rtype: bool
	"""
	table = [[False] * (len(s)+1) for _ in range(len(p)+1)]
	table[0][0] = True
		
	for i in range(2, len(p) + 1):
		if p[i-1] == "*":
			table[i][0] = table[i - 2][0]
	
	for i in range(1, len(p) + 1):
		for j in range(1, len(s) + 1):
			if p[i - 1] != "*":
				# Update the table by referring the diagonal element.
				if (p[i - 1] == s[j - 1] or p[i - 1] == '.'):
					table[i][j] = table[i - 1][j - 1]
			else:
				# Eliminations (referring to the vertical element)
				# Either refer to the one before previous or the previous.
				# I.e. * eliminate the previous or count the previous.
				table[i][j] = table[i - 2][j] or table[i - 1][j]

				# Propagations (referring to the horizontal element)
				# If p's previous one is equal to the current s, with
				# helps of *, the status can be propagated from the left.
				if p[i - 2] == s[j - 1] or p[i - 2] == '.':
					table[i][j] = table[i][j - 1] | table[i][j]
	
	return table[-1][-1]
	
s = "abbc"
p = "ab*c"
print(isMatch(s, p))