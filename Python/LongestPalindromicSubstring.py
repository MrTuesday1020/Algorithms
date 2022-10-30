def longestPalindrome(s):
	"""
	:type s: str
	:rtype: str
	"""
	res = ""
	for i in range(len(s)):
		temp = helper(i,i,s)
		if len(temp) > len(res):
			res = temp
		temp = helper(i,i+1,s)
		if len(temp) > len(res):
			res = temp
	
	return res

def helper(l,r,s):
	while l >= 0 and r < len(s) and s[l] == s[r]:
		l -= 1
		r += 1
	return s[l+1:r]
	
print(longestPalindrome("aacdefcaa"))