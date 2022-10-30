def combinationSum2(candidates: 'List[int]', target: 'int') -> 'List[List[int]]':
	if not candidates:
		return
		
	ans = []
	candidates.sort()
	dfs(candidates, target, [], 0, ans)
	return ans
	
def dfs(candidates, target, path, index, ans):
	if target < 0:
		return
	elif target == 0 and path not in ans:
		ans.append(path)
	else:
		for i in range(index, len(candidates)):
			dfs(candidates, target-candidates[i], path+[candidates[i]], i+1, ans)


candidates = [4,4,2,1,4,2,2,1,3]
target = 6
print(combinationSum2(candidates, target))