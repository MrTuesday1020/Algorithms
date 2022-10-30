"""
https://leetcode.com/problems/letter-combinations-of-a-phone-number/
"""

def letterCombinations(digits):
	"""
	:type digits: str
	:rtype: List[str]
	"""
	
	numbers = {
		2:["a","b","c"],
		3:["d","e","f"],  
		4:["g","h","i"],  
		5:["j","k","l"],  
		6:["m","n","o"],  
		7:["p","r","q","s"],  
		8:["t","u","v"],  
		9:["w","x","y","z"]
	}
		
	combs = []
	
	while digits:
		current_number = int(digits[0])
		digits = digits[1:]
		if not combs:
			for char in numbers[current_number]:
				combs.append(char)
		else:
			temp = []
			for comb in combs:
				for char in numbers[current_number]:
					temp.append(comb+char)
			combs = temp
	
	return combs

print(letterCombinations("23"))