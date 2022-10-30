def removeDuplicates(nums):
	if not nums:
		return
		
	count = 1 # count length
	_count = 1 # count duplicates
	for i in range(1, len(nums)):
		if nums[i] == nums[i-1]:
			if _count == 1:
				nums[count] = nums[i]
				count += 1	
			_count += 1		
		elif nums[i] != nums[i-1]:
			nums[count] = nums[i]
			_count = 1
			count += 1			

	return count		
	
nums = [0,0,1,1,1,2,2,2,2,2,2,3,3,4]
print(removeDuplicates(nums))