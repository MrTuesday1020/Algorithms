def QuickSort1(arr):
	if len(arr) < 2:
		return arr
	else:
		pivot = arr[0]
		left = [i for i in arr if i < pivot]
		right = [i for i in arr if i > pivot]
		return QuickSort1(left) + [pivot] + QuickSort1(right)
		
print(QuickSort1([1,2,5,3,4,7,8,6]))


# This function takes last element as pivot, places  the pivot element at its correct position in sorted 
# array, and places all smaller (smaller than pivot) to left of pivot and all greater elements to right of pivot 
def partition(arr, low, high): 
	i =  low - 1	# index of smaller element 
	pivot = arr[high]	 # pivot 

	for j in range(low , high): 

		# If current element is smaller than or 
		# equal to pivot 
		if arr[j] <= pivot: 
		
			# increment index of smaller element 
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i] 

	arr[i+1], arr[high] = arr[high], arr[i+1] 
	return i + 1

# The main function that implements QuickSort 
# arr[] --> Array to be sorted, 
# low --> Starting index, 
# high --> Ending index 

# Function to do Quick sort 
def QuickSort2(arr,low,high): 
	if low < high: 

		# pi is partitioning index, arr[p] is now 
		# at right place 
		pi = partition(arr,low,high) 

		# Separately sort elements before 
		# partition and after partition 
		QuickSort2(arr, low, pi-1) 
		QuickSort2(arr, pi+1, high) 

# Driver code to test above 
arr = [10, 7, 8, 9, 1, 5] 
n = len(arr) 
QuickSort2(arr,0,n-1) 
print(arr)