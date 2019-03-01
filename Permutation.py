"""
https://leetcode.com/problems/permutations/
"""

def permutation1(arr):
	if len(arr) == 0:
		return []
	elif len(arr) == 1:
		return [arr]
	else:
		res = []
		for i in range(len(arr)):
			first_number = arr[i]
			rest_numbers = arr[:i] + arr[i+1:]
			for item in permutation1(rest_numbers):
				res.append([first_number] + item)
		return res
		
def permutation2(arr):
	if len(arr) == 0:
		yield []
	elif len(arr) == 1:
		yield arr
	else:
		for i in range(len(arr)):
			first_number = arr[i]
			rest_numbers = arr[:i] + arr[i+1:]
			for item in permutation1(rest_numbers):
				yield [first_number] + item
		
arr = list("abcdefgh")
for item in permutation1(arr):
	print(item)