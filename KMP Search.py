"""
http://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
https://www.geeksforgeeks.org/kmp-algorithm-for-pattern-searching/
"""

def KMPSearch(pat, txt): 
	"""
	:type pat: str
	:type txt: str
	"""
	M = len(pat) 
	N = len(txt) 
	# create lps[] that will hold the longest prefix suffix  
	# values for pattern 
	lps = [0]*M 
	j = 0 # index for pat[]
	# Preprocess the pattern (calculate lps[] array) 
	computeLPS(pat, M, lps)
	
	i = 0 # index for txt[] 
	while i < N: 
		# character matches
		if pat[j] == txt[i]: 
			i += 1
			j += 1
		# pattern matches
		if j == M: 
			print("Found pattern at index " + str(i-j))
			j = lps[j-1] 
		# mismatch after j matches 
		elif i < N and pat[j] != txt[i]: 
			# Do not match lps[0..lps[j-1]] characters, 
			# they will match anyway 
			if j != 0: 
				j = lps[j-1] 
			else: 
				i += 1
				
def computeLPS(pat, M, lps): 
	"""
	:type pat: str
	:type M: int
	:type lps: List[int]
	"""
	length = 0 # length of the previous longest prefix suffix 
	lps[0] # lps[0] is always 0 
	i = 1 # current index
	# the loop calculates lps[i] for i = 1 to M-1 
	while i < M: 
		if pat[i]== pat[length]: 
			length += 1
			lps[i] = length
			i += 1
		else: 
			# to search step. 
			if length != 0: 
				length = lps[length-1] 
				# Also, note that we do not increment i here 
			else: 
				lps[i] = 0
				i += 1
	
	print(lps)
				
txt = "ABABDABACDABABCABAB"
pat = "ABABCABAB"
KMPSearch(pat, txt) 