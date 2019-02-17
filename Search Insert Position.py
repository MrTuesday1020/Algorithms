"""
https://leetcode.com/problems/search-insert-position/
"""

def searchInsert(nums: 'List[int]', target: 'int') -> 'int':
	if not nums:
		return 0
	
	if target <= nums[0]:
		return 0
	
	if target > nums[-1]:
		return len(nums)
		
	start = 0
	end = len(nums)-1
	
	while start < end:
		mid = (start + end) // 2
		if nums[mid] == target:
			return mid
		elif target <= nums[mid]:
			end = mid
		else:
			start = mid + 1
			
	return start
	
nums = [1,3,5,6]
target = 4
print(searchInsert(nums, target))