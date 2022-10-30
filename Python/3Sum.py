"""
https://leetcode.com/problems/3sum/
"""

def threeSum(nums):
	"""
	:type nums: List[int]
	:rtype: List[List[int]]
	"""
	if len(nums) < 3:
		return[]
	nums = sorted(nums)
	res = []
	for i in range(len(nums)):
		# remove redundance
		if i > 0 and nums[i] == nums[i-1]:
			continue
		
		target = -1 * nums[i]
		left = i + 1
		right = len(nums) - 1
		# scan the array from left and right sides at the same time
		while left < right:
			if nums[left] + nums[right] < target:
				left += 1
			elif nums[left] + nums[right] > target:
				right -= 1
			else:
				res.append([nums[i], nums[left], nums[right]])
				left += 1
				# remove redundance
				while left < right and nums[left] == nums[left-1]:
					left += 1
					
	return res
	
print(threeSum([-1, 0, 1, 2, -1, -4]))