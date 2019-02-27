"""
https://leetcode.com/problems/first-missing-positive/
"""

def firstMissingPositive(nums):
	n = len(nums)
	
	for i in range(n):
		while 0 <= nums[i]-1 < len(nums) and nums[nums[i]-1] != nums[i]:
			tmp = nums[i]-1
			nums[i], nums[tmp] = nums[tmp], nums[i]
			
	for i in range(n):
		if nums[i] != i+1:
			return i+1
			
	return n+1
	
nums = [3,4,-1,1]
print(firstMissingPositive(nums))