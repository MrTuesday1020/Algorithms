"""
https://leetcode.com/problems/remove-element/
"""

def removeElement(nums, val):
	"""
	:type nums: List[int]
	:type val: int
	:rtype: int
	"""
	if not nums:
		return
		
	cnt = 0
	
	for i in range(len(nums)):
		if nums[i] != val:
			nums[cnt] = nums[i]
			cnt += 1
		
	return nums
	
print(removeElement([0,1,2,2,3,0,4,2], 2))