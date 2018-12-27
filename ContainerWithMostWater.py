"""
https://leetcode.com/problems/container-with-most-water/
"""

def maxArea(height):
	"""
	:type height: List[int]
	:rtype: int
	"""
	maxarea = 0
	while(height):
		if height[0] < height[-1]:
			area = (len(height) - 1) * height.pop(0)
		else:
			area = (len(height) - 1) * height.pop()
		maxarea = max(area, maxarea)
	
	return maxarea
			
height = [1,8,6,2,5,4,8,3,7]
print(maxArea(height))