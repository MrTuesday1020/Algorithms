"""
https://leetcode.com/problems/jump-game/
"""

def canJump(nums):
	n = len(nums)
	end = 0
	for i in range(n):
		if i > end:
			return False
		end = max(end, i + nums[i])
	return True
	
print(canJump([2,3,1,1,4]))