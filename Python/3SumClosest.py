"""
https://leetcode.com/problems/3sum-closest/
"""

import math

def threeSumClosest(nums, target):
	"""
	:type nums: List[int]
	:type target: int
	:rtype: int
	"""
	
	if len(nums) < 3:
		return
	nums = sorted(nums)
	current_diff = float('inf')
	res = target
	for i in range(len(nums)):
		# remove redundance
		if i > 0 and nums[i] == nums[i-1]:
			continue
		
		temp_target = target - nums[i]
		left = i + 1
		right = len(nums) - 1
		# scan the array from left and right sides at the same time
		while left < right:
			if nums[left] + nums[right] < temp_target:
				temp = math.fabs(target - nums[i] - nums[left] - nums[right])
				if temp < current_diff:
					current_diff = temp
					res = nums[i] + nums[left] + nums[right]
				left += 1
			elif nums[left] + nums[right] > temp_target:
				temp = math.fabs(target - nums[i] - nums[left] - nums[right])
				if temp < current_diff:
					current_diff = temp
					res = nums[i] + nums[left] + nums[right]
				right -= 1
			else:
				return target
					
	return res
	
print(threeSumClosest([-1, 2, 1, -4],1))