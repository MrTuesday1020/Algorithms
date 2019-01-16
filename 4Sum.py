"""
https://leetcode.com/problems/4sum/
"""

def fourSum(nums, target):
	"""
	:type nums: List[int]
	:type target: int
	:rtype: List[List[int]]
	"""
	
	if len(nums) < 4:
		return[]
	nums = sorted(nums)
	
	res = []
	
	for i in range(len(nums)-1):
		# remove redundance
		if i > 0 and nums[i] == nums[i-1]:
			continue
		
		for j in range(i+1,len(nums)):
			if j > i + 1 and nums[j] == nums[j-1]:
				continue
			
			temp_target = target - nums[i] - nums[j]
			left = j + 1
			right = len(nums) - 1
			
			while left < right:
				if nums[left] + nums[right] < temp_target:
					left += 1
				elif nums[left] + nums[right] > temp_target:
					right -= 1
				else:
					res.append([nums[i], nums[j], nums[left], nums[right]])
					left += 1
					# remove redundance
					while left < right and nums[left] == nums[left-1]:
						left += 1
						
	return res
	
print(fourSum([0,0,0,0],0))