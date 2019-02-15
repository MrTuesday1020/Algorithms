"""
https://leetcode.com/problems/search-in-rotated-sorted-array/
https://leetcode.windliang.cc/leetCode-33-Search-in-Rotated-Sorted-Array.html
这个算法基于一个事实，数组从任意位置劈开后，至少有一半是有序的，什么意思呢？
比如 [ 4 5 6 7 1 2 3] ，从 7 劈开，左边是 [ 4 5 6 7] 右边是 [ 7 1 2 3]，左边是有序的。基于这个事实。
我们可以先找到哪一段是有序的 (只要判断端点即可)，然后看 target 在不在这一段里，如果在，那么就把另一半丢弃。如果不在，那么就把这一段丢弃。
"""

def search(self, nums: 'List[int]', target: 'int') -> 'int':
	start = 0
	end = len(nums) - 1

	while(start <= end):
		mid = (start + end) // 2
		if nums[mid] == target:
			return mid
			
		if nums[start] <= nums[mid]:
			if (target >= nums[start] and target < nums[mid]):
				end = mid - 1
			else:
				start = mid + 1
		else:
			if (target > nums[mid] and target <= nums[end]):
				start = mid + 1
			else:
				end = mid - 1
				
	return -1