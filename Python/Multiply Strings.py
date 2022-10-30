"""
https://leetcode.com/problems/multiply-strings/
"""

def multiply(num1, num2):
	ans = 0
	offset2 = 1
	for i in reversed(num1):
		offset1 = 1
		tmp = 0
		for j in reversed(num2):
			tmp += int(i) * int(j) * offset1
			offset1 *= 10
		ans += tmp * offset2
		offset2 *= 10
	
	return str(ans)
			
num1 = "123"
num2 = "456"
print(multiply(num1, num2))