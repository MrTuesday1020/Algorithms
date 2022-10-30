package main

import (
	"fmt"
)

//func nextGreaterElement(nums1 []int, nums2 []int) []int {
//	ret := make([]int, len(nums1))
//	for i := 0; i < len(nums1); i++ {
//		got := false
//		equal := false
//		for j := 0; j < len(nums2); j++ {
//			if nums1[i] == nums2[j] {
//				// find equal element
//				equal = true
//			}
//			if nums1[i] < nums2[j] && equal {
//				// find next greater element
//				ret[i] = nums2[j]
//				got = true
//				break
//			}
//		}
//		if !got {
//			ret[i] = -1
//		}
//	}
//	return ret
//}

// 下面这种时间复杂度低 但是空间复杂度高
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{nums2[0]}
	m := map[int]int{}
	for i := 1; i < len(nums2); i++ {
		if nums2[i] < stack[len(stack)-1] {
			// current number smaller than top element in stack
			stack = append(stack, nums2[i])
			continue
		}
		// pop the element larger than the current number
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			m[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	// left numbers in stack cant find next greater element 
	for _, element := range stack {
		m[element] = -1
	}
	
	ret := make([]int, len(nums1))
	for i, num := range nums1 {
		ret[i] = m[num]
	}
	return ret
}

func main() {
	fmt.Println(nextGreaterElement([]int{4,1,2}, []int{1,3,4,2}))
}