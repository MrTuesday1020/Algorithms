"""
https://leetcode.com/problems/zigzag-conversion/

寻找数学规律
"""

def convert(s, numRows):
	"""
	:type s: str
	:type numRows: int
	:rtype: str
	"""
	if numRows == 0:
		return
	if numRows == 1:
		return s
		
	res = ""
	steps = 2 * (numRows - 1)
	for i in range(numRows):
		index = i
		while index < len(s):
			res += s[index]
			mid_index = index + steps - 2*i
			if i != 0 and i != numRows-1 and mid_index < len(s):
				res += s[mid_index]
			index += steps
			
	return res
	
print(convert("PAYPALISHIRING", 4))