"""
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
"""

def search(nums, target):
	start = 0
	end = len(nums) - 1

	while(start <= end):
		mid = (start + end) // 2
		if nums[mid] == target:
			return True
		while start < mid and nums[start] == nums[mid]: # tricky part
			start += 1
			
		# the first half is ordered
		if nums[start] <= nums[mid]:
			if nums[start] <= target < nums[mid]:
				end = mid - 1
			else:
				start = mid + 1
		else:
			if nums[mid] < target <= nums[end]:
				start = mid + 1
			else:
				end = mid - 1
				
	return False