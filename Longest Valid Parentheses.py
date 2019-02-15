"""
https://leetcode.com/problems/longest-valid-parentheses/
"""

def longestValidParentheses(s: 'str') -> 'int':
	"""
	Instead of finding every possible string and checking its validity, we can make use of stack while scanning the given string to check if the string scanned so far is valid, and also the length of the longest valid string. In order to do so, we start by pushing -1−1 onto the stack.

	For every ‘(’ encountered, we push its index onto the stack.

	For every ‘)’ encountered, we pop the topmost element and subtract the current element's index from the top element of the stack, which gives the length of the currently encountered valid string of parentheses. If while popping the element, the stack becomes empty, we push the current element's index onto the stack. In this way, we keep on calculating the lengths of the valid substrings, and return the length of the longest valid string at the end.
	"""
	
	if not s:
		return 0

	ans = 0
	stack = [-1]

	for i in range(len(s)):
		if s[i] == '(':
			stack.append(i)
		else:
			stack.pop()
			if not stack:
				stack.append(i)
			else:
				ans = max(ans,i - stack[-1])
	return ans
	
print(longestValidParentheses("()(()"))