"""
https://leetcode.com/problems/trapping-rain-water/
https://leetcode.windliang.cc/leetCode-42-Trapping-Rain-Water.html
"""

def trap(height):
	ans = 0
	n = len(height)
	max_left = [0] * n
	max_right = [0] * n
	
	for i in range(1, n):
		max_left[i] = max(max_left[i-1], height[i-1])
		max_right[n-i-1] = max(max_right[n-i], height[n-i])

	for i in range(n):
		min_side = min(max_left[i], max_right[i])
		if height[i] < min_side:
			ans += min_side - height[i]
			
	return ans
	
height = [0,1,0,2,1,0,1,3,2,1,2,1]
print(trap(height))