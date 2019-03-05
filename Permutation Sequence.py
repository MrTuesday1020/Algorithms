"""
https://leetcode.com/problems/permutation-sequence/
"""

def getPermutation(n, k):
	import math
	nums = [i for i in range(1,n+1)]
	ans = ""
	
	# arr[k-1] means the kth element
	k -= 1
	while n > 0:
		n -= 1
		# get the index of current digit
		index, k = divmod(k, math.factorial(n))
		print(index, k)
		ans += str(nums[index])
		# remove handled number
		nums.remove(nums[index])
		
	return ans
		
print(getPermutation(3, 3))