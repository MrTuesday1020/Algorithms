"""
https://leetcode.com/problems/remove-duplicates-from-sorted-array/
"""

def removeDuplicates(nums):
	"""
	:type nums: List[int]
	:rtype: int
	"""
	if not nums:
		return
		
	count = 1
	for i in range(1,len(nums)):
		if nums[i] != nums[i-1]:
			nums[count] = nums[i]
			count += 1
	
	return count
		
print(removeDuplicates([0,0,1,1,1,2,2,3,3,4]))