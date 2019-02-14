"""
https://leetcode.com/problems/substring-with-concatenation-of-all-words/
"""

def findSubstring(s, words):
	"""
	:type s: str
	:type words: List[str]
	:rtype: List[int]
	"""
	if not s:
		return []
		
	if not words:
		return []
	
	len_of_word = len(words[0])
	no_of_words = len(words)
	len_of_substring = len_of_word * no_of_words
	len_of_string = len(s)
	
	if len_of_string - len_of_substring < 0:
		return []
	
	sorted_words = sorted(words)
	
	ans = []
	for i in range(0, len_of_string - len_of_substring + 1):
		# construct 
		substring = s[i:i+len_of_substring]
		words_of_substring = []
		for j in range(no_of_words):
			current_word = substring[j*len_of_word:(j+1)*len_of_word]
			if current_word not in words:
				# if this word not in words list, then we don't need to consider the words after it
				break
			else:
				words_of_substring.append(current_word)
		sorted_words_of_substring = sorted(words_of_substring)
		# we compare the sorted words list
		if sorted_words_of_substring == sorted_words:
			ans.append(i)
			
	return ans

s = "barfoothefoobarman"
words = ["foo","bar"]
print(findSubstring(s, words))