"""
	https://www.youtube.com/watch?v=dQWsEWgbgc8&list=PLAnjpYDY-l8IacYv_2lIZxNrQmkY3paSN&index=5
	给定一个数组 {3, 1, 2, 1} 和一个数字k =4。求这个数组的一个最长连续子数组，这个最长连续子数组中所有数字的必须小于或等于k。
	例如，上面这个例子中，连续子数组有这么多种情况：
	{3}, {1}, {2}, {1}, {3, 1}, {1, 2}, {2, 1}, {3, 1, 2}, {1, 2, 1}, {3, 1, 2, 1}。
	其中符合条件的就只有{1, 2, 1}
"""

def func(array, k):
	if not array or k <=0:
		return
	
	# init
	current_sum = array[0]
	current_length = 1
	sums = []
	sums.append(current_sum)
	length = []
	length.append(current_length)
	
	for i in range(1, len(array)):
		current_length += 1
		current_sum = current_sum + array[i]
		if current_sum > k:
			current_length -= 1
			current_sum = current_sum - array[i-current_length]
			sums.append(current_sum)
			length.append(current_length)
		else:
			sums.append(current_sum)
			length.append(current_length)
	
	print(sums)
	print(length)
	
	max_length = 0
	for i in range(len(array)):
		if sums[i] <= k:
			max_length = max(max_length, length[i])

	return max_length	
	
#array = [1,2,1,3,1,1,1]
#k = 4
array = [1,1,1,9,9,1,1,1,1]
k = 4
max_length = func(array, k)

print("The max length is",max_length)