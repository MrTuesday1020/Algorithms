"""
https://leetcode.com/problems/powx-n/
"""

def myPow(x, n):
	ans = 1
	if n == 0:
		return ans
	elif n < 0:
		n = -n
		x = 1/x
			
	while n:
		if n & 1:
			ans *= x
		x *= x
		n >>= 1
		
	return ans
	
print(myPow(2, -2))