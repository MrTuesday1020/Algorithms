"""
https://leetcode.com/problems/valid-number/
https://leetcode.com/problems/valid-number/discuss/23725/C%2B%2B-My-thought-with-DFA
"""

def isNumber(s):
	s = s.strip()
	state = 0
	flag = 0
	
	for char in s:
		if ord(char) >= 48 and ord(char) <= 57:
			flag = 1
			if state <= 2:
				state = 2
			else:
				state = 5 if state <= 5 else 7
		elif char == "+" or char == "-":
			if state==0 or state==3:
				state += 1
			else:
				return False
		elif char == ".":
			if state <= 2:
				state = 6;
			else:
				return False
		elif char == 'e':
			if flag and ( state == 2 or state == 6 or state == 7 ):
				state = 3
			else:
				return False
		else:
			return False
	
	return state == 2 or state == 5 or (flag and state == 6) or state == 7
			
	
s = " 0.1 "
print(isNumber(s))