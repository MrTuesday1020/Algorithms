"""
https://leetcode.com/problems/implement-strstr/
"""

def strStr(haystack, needle):
	"""
	:type haystack: str
	:type needle: str
	:rtype: int
	"""
	def BadCharHeuristic(pattern):
		badchar = [-1] * 256
		# record the last occurance of the character
		for i in range(len(pattern)):
			badchar[ord(pattern[i])] = i
			
		return badchar
	
	def BMSearch(text, pattern):
		badchar = BadCharHeuristic(pattern)
		m = len(text)
		n = len(pattern)
		
		i = 0 # current index
		while(i <= m-n):
			j = n-1
			
			while j >= 0 and pattern[j] == text[i+j]:
				j -= 1
				
			if j < 0:
				# find matched pattern
				return i
			else:
				# move pattern
				i += max(1, j-badchar[ord(text[i+j])])
		return -1
		
	return BMSearch(haystack, needle)
	
def main():
	haystack = "mississippi"
	needle = "pi"
	print(strStr(haystack, needle))
	
if __name__ == "__main__":
	main()