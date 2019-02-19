def combinationSum(candidates: 'List[int]', target: 'int') -> 'List[List[int]]':
	if not candidates:
		return []
		
	ans = []
	length = len(candidates)
	dfs(candidates, target, [], 0, ans, length)
	return ans

def dfs(candidates, target, path, index, ans, length):
	if target < 0:
		return
	elif target == 0:
		ans.append(path)
	else:
		for i in range(index, length):
			dfs(candidates, target-candidates[i], path+[candidates[i]], i, ans, length)
	
	
candidates = [2,3,5]
target = 8
print(combinationSum(candidates, target))