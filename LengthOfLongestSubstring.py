def lengthOfLongestSubstring(s):
		"""
		:type s: str
		:rtype: int
		"""
		if len(s) == 0:
			return 0
		
		mapping = {}
		start = 0
		maxlen = 0
		
		for i in range(len(s)):
			if s[i] in mapping:
				if mapping[s[i]] >= start:
					start = mapping[s[i]] + 1
			mapping[s[i]] = i
			maxlen = max(maxlen, i-start+1)
		
		return maxlen
		
print(lengthOfLongestSubstring("dvdf"))