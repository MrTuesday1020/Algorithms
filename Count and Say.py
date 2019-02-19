"""
https://leetcode.com/problems/count-and-say/
"""

def countAndSay(n: 'int') -> 'str':
	ans = "1"
	
	for i in range(2, n+1):
		copy = ""
		char = ans[0]
		cnt = 0
		for this_char in ans:
			if char == this_char:
				cnt += 1
			else:
				copy += str(cnt) + char
				cnt = 1
				char = this_char
		copy += str(cnt) + char
		ans = copy
		
	return ans
	
print(countAndSay(6))