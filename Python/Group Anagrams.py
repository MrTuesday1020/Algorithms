"""
https://leetcode.com/problems/group-anagrams/
"""

def groupAnagrams(strs):

	d = {}
	
	for s in strs:
		key = tuple(sorted(s))
		if key in d:
			d[key].append(s)
		else:
			d[key] = [s]
			
	ans = []
	for value in d.values():
		ans.append(value)		
	
	return ans
	
print(groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]))