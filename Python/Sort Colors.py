"""
https://leetcode.com/problems/sort-colors/
"""

def sortColors(nums):
	"""
	[, i): 0 
	[i, j]: 1
	(k, ...]: 2
	Once j meets k, the sorting is complete
	"""
	
	i,j,k = 0,0,len(nums)-1
	
	while j <= k:
		if nums[j] == 0:
			nums[i], nums[j] = nums[j], nums[i]
			i += 1
			j += 1
		elif nums[j] == 1:
			j += 1
		else:
			nums[j], nums[k] = nums[k], nums[j]
			k -= 1