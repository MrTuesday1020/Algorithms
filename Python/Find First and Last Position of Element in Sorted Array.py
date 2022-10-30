"""
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
"""

def searchRange(nums: 'List[int]', target: 'int') -> 'List[int]':
	start = -1
	end = -1
	if not nums:
		return [start, end]
		
	if nums[0] > target or nums[-1] < target:
		return [start, end]
		
	flag = 0 # flag为0表示还没找到，为1表示已经找到
	for i in range(len(nums)):
		if nums[i] == target and flag == 0:
			start = i
			end = i
			flag = 1
		elif nums[i] == target and flag == 1:
			end = i
		elif nums[i] > target:
			break
			
	return [start, end]
	
nums = [5,7,7,8,8,10]
print(searchRange(nums, 6))