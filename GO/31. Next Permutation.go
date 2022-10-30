package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int)  {
	if len(nums) < 2 {
		return
	}
	
	leftID, rightID := 0, 0
	got := false
	
	// find the left num, which is the first num in desc order from end
	// and set next num as the first right num 
	for i := len(nums)-2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			got = true
			leftID = i
			rightID = i+1
			break
		}
	}
			
	if !got {
		// nums are in desc order
		reverse(nums)
		return
	}
		
	// find the smallest num on the right side of the left num
	for i := rightID; i < len(nums); i++ {
		if nums[i] < nums[rightID] && nums[i] > nums[leftID]{
			rightID = i
		}
	}
			
	nums[leftID], nums[rightID] = nums[rightID], nums[leftID]
		
	sort.Ints(nums[leftID+1:])
}

func reverse(nums []int) {
	for i := 0; i < len(nums) / 2; i++ {
		nums[i], nums[len(nums) -i - 1] = nums[len(nums) - i -1 ], nums[i]
	}
}

func main() {
	arr := []int{1,2,3}
	nextPermutation(arr)
	fmt.Println(arr)
	
	arr = []int{1,3,2}
	nextPermutation(arr)
	fmt.Println(arr)
	
	arr = []int{3,2,1}
	nextPermutation(arr)
	fmt.Println(arr)
	
	arr = []int{2,3,1}
	nextPermutation(arr)
	fmt.Println(arr)
	
	arr = []int{2,3,1,3,3}
	nextPermutation(arr)
	fmt.Println(arr)
}