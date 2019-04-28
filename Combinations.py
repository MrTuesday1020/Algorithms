"""
https://leetcode.com/problems/combinations/
"""

def combine(n, k):
	ans = []
	nums = [i+1 for i in range(n)]
	
	dfs(ans, 0, [], nums, k)
	return ans
	
def dfs(ans, start, path, nums, k):
	if k == 0:
		ans.append(path)
		return # backtracking 
	for i in range(start, len(nums)):
		dfs(ans, i+1, path+[nums[i]], nums, k-1)
	
print(combine(4, 2))