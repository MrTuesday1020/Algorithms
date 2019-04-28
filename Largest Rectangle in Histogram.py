def largestRectangleArea(heights):
	n = len(heights)
	lowerFromLeft = [-1] * n
	lowerFromRight = [n] * n
	
	for i in range(1, n):
		p = i-1
		while p >= 0 and heights[p] >= heights[i]:
			p = lowerFromLeft[p]
		lowerFromLeft[i] = p
			
	for i in range(n-2, -1, -1):
		p = i+1
		while p < n and heights[p] >= heights[i]:
			p = lowerFromRight[p]
		lowerFromRight[i] = p
		
	maxarea = 0
	for i in range(n):
		maxarea = max(maxarea, heights[i] * (lowerFromRight[i] - lowerFromLeft[i] - 1))
		
	return maxarea
	
heights = [2,1,5,6,2,3]
print(largestRectangleArea(heights))