"""
https://leetcode.com/problems/next-permutation/
"""

def nextPermutation(nums):
	"""
	:type nums: List[int]
	:rtype: void Do not return anything, modify nums in-place instead.
	"""
	
	if not nums:
		return
	
	node1 = None
	node2 = None
	
	for i in range(len(nums)-1, 0, -1):
		if nums[i-1] < nums[i]:
			node1 = i-1
			break
			
	if node1 != None:
		for i in range(node1+1,len(nums)):
			if nums[i] > nums[node1]:
				node2 = i
			
	if node1 != None and node2 != None:
		nums[node1], nums[node2] = nums[node2], nums[node1]
		nums[node1+1:] = sorted(nums[node1+1:])
		print(nums)
	else:
		#nums = nums.reverse()
		nums = sorted(nums)
		print(nums)
			
			
nums = [3,2,1]
nextPermutation(nums)