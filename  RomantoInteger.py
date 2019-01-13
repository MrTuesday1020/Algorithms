"""
https://leetcode.com/problems/roman-to-integer/
"""

def romanToInt(s):
	res = 0
	
	while s:
		if s[0] == "M":
			res += 1000
			s = s[1:]
		elif s[0] == "D":
			res += 500
			s = s[1:]
		elif s[0] == "C":
			if len(s) > 1 and s[1] == "D":
				res += 400
				s = s[2:]
			elif len(s) > 1 and s[1] == "M":
				res += 900
				s = s[2:]
			else:
				res += 100
				s = s[1:]
		elif s[0] == "L":
			res += 50
			s = s[1:]
		elif s[0] == "X":
			if len(s) > 1 and s[1] == "L":
				res += 40
				s = s[2:]
			elif len(s) > 1 and s[1] == "C":
				res += 90
				s = s[2:]
			else:
				res += 10
				s = s[1:]
		elif s[0] == "V":
			res += 5
			s = s[1:]
		elif s[0] == "I":
			if len(s) > 1 and s[1] == "V":
				res += 4
				s = s[2:]
			elif len(s) > 1 and s[1] == "X":
				res += 9
				s = s[2:]
			else:
				res += 1
				s = s[1:]
		else:
			return "Invalid string!"
			
	return res
		
print(romanToInt("LVIII"))