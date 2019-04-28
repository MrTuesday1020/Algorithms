"""
https://leetcode.com/problems/subsets/
"""

def subsets(nums):
	res = [[]]
	for num in nums:
		res += [item+[num] for item in res]
	return res
	
print(subsets([1,2,3]))