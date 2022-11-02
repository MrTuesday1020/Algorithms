"""
https://leetcode.com/problems/minimum-window-substring/solution/
"""

def minWindow(s, t):
	import collections
	need = collections.Counter(t)            #hash table to store char frequency
	missing = len(t)                         #total number of chars we care
	start, end = 0, 0
	i = 0
	for j, char in enumerate(s, 1):          # enumerate j form 1,2,3... and element one by one
		if need[char] > 0:
			missing -= 1
		# if char is not in need, the value will be -1
		need[char] -= 1
		
		if missing == 0:                     #match all chars
			# 右移到不能再移动为止
			while i < j and need[s[i]] < 0:  #remove chars to find the real start
				need[s[i]] += 1
				i += 1
			need[s[i]] += 1                  #make sure the first appearing char satisfies need[char]>0
			missing += 1                     #we missed this first char, so add missing by 1
			if end == 0 or j-i < end-start:  #update window
				start, end = i, j
			i += 1                           #update i to start+1 for next window
	return s[start:end]
	
S = "ADOBECODEBANC"
T = "ABC"
print(minWindow(S, T))