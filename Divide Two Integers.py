"""
https://leetcode.com/problems/divide-two-integers/
"""

def divide(dividend, divisor):
	"""
	:type dividend: int
	:type divisor: int
	:rtype: int
	"""
	if dividend == 0:
		return 0
		
	mark = (dividend > 0) ^ (divisor > 0)
	dividend, divisor = abs(dividend), abs(divisor)
		
	cnt = 0
	while dividend >= divisor:
		x = divisor
		i = 1
		while dividend >= x + x:
			x += x
			i += i
		dividend -= x
		cnt += i
				
	cnt = 2**31-1 if not mark and cnt > 2**31-1 else cnt
		
	return -cnt if mark else cnt
		
dividend = -1
divisor = -1
print(divide(dividend, divisor))