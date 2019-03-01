"""
https://leetcode.com/problems/jump-game-ii/
"""

def jump(nums):
	if not nums:
		return
	if len(nums) == 1:
		return 0
	
	end, maxPosition, step = 0, 0, 0
	for i in range(len(nums)-1):
		maxPosition = max(maxPosition, nums[i] + i)
		if i == end:
			end = maxPosition
			step += 1
	
	return step
	
nums = [2,3,1,1,4]
print(jump(nums))