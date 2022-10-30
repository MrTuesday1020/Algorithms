"""
https://leetcode.com/problems/longest-common-prefix/
"""

def longestCommonPrefix(strs):
	"""
	:type strs: List[str]
	:rtype: str
	"""
	if not strs:
		return ""
		
	
	def LCP(strs, left, right):
		if left == right:
			return strs[left]
		else:
			mid = (left + right) // 2
			lres = LCP(strs, left, mid)
			rres = LCP(strs, mid+1, right)
			
			shorter = lres if len(lres) < len(rres) else rres
			for i in range(len(shorter)):
				if lres[i] != rres[i]:
					return lres[0:i]
			return shorter
					
	return LCP(strs, 0, len(strs)-1)
	
print(longestCommonPrefix(["dog","racecar","car"]))