"""
https://leetcode.com/problems/valid-parentheses/
"""

def isValid(s):
	"""
	:type s: str
	:rtype: bool
	"""
	
	queue = []
	for parenthes in s:
		if len(queue) == 0:
			queue.append(parenthes)
		elif queue[-1] == '(' and parenthes == ')':
			queue.pop()
		elif queue[-1] == '{' and parenthes == '}':
			queue.pop()
		elif queue[-1] == '[' and parenthes == ']':
			queue.pop()
		else:
			queue.append(parenthes)
	
	if len(queue) == 0:
		return True
	else:
		return False
		
print(isValid("{[]}"))