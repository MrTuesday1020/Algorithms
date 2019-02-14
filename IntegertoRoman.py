"""
https://leetcode.com/problems/integer-to-roman/
"""



def intToRoman(num):
	"""
	:type num: int
	:rtype: str
	"""
	string1 = ["","I","II","III","IV","V","VI","VII","VIII","IX"]
	string10 = ["","X","XX","XXX","XL","L","LX","LXX","LXXX","XC"]
	string100 = ["","C","CC","CCC","CD","D","DC","DCC","DCCC","CM"]
	string1000 = ["","M","MM","MMM"]
	
	res = string1000[num//1000] + string100[(num%1000)//100] + string10[(num%100)//10] + string1[num%10]
	
	return res
	
print(intToRoman(58))