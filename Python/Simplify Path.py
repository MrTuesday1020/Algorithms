"""
https://leetcode.com/problems/simplify-path/
"""

def simplifyPath(path):
	stack = []
	
	for token in path.split('/'):
		if token in ('', '.'):
			pass
		elif token == '..':
			if stack: stack.pop()
		else:
			stack.append(token)
	return '/' + '/'.join(stack)
	
path = "/./"
print(simplifyPath(path))