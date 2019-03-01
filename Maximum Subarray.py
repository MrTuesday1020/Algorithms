"""
https://leetcode.com/problems/maximum-subarray/
"""

def maxSubArray(nums):
	ans = nums[0]
	opt = [nums[0]] * len(nums)
	
	for i in range(1, len(nums)):
		if opt[i-1] > 0:
			opt[i] = nums[i] + opt[i-1]
		else:
			opt[i] = nums[i]
			
		ans = max(opt[i], ans)
		
	return ans
	
print(maxSubArray([-2,1,-3,4,-1,2,1,-5,4]))
		