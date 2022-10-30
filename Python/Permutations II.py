"""
https://leetcode.com/problems/permutations-ii/
"""

def permuteUnique(nums):
	if len(nums) == 0:
		return []
	elif len(nums) == 1:
		return [nums]
	else:
		res = []
		seen = set()
		for i in range(len(nums)):
			if nums[i] not in seen:
				seen.add(nums[i])
				first_number = nums[i]
				rest_numbers = nums[:i] + nums[i+1:]
				for item in permuteUnique(rest_numbers):
					res.append([first_number] + item)
		return res

nums = [1,1,3]
print(permuteUnique(nums))