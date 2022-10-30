package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	l1 := len(nums1)
	l2 := len(nums2)
	nums := make([]int, l1+l2)
	for i + j < (l1 + l2) / 2 + 1 {
		if i < l1 && j < l2 {
			if nums1[i] < nums2[j] {
				nums[i+j] = nums1[i]
				i += 1
			} else {
				nums[i+j] = nums2[j]
				j += 1
			}
		} else if i < l1 {
			nums[i+j] = nums1[i]
			i += 1
		} else {
			nums[i+j] = nums2[j]
			j += 1
		}
	}
	fmt.Println(nums)
	
	if (l1 + l2) % 2 == 0 {
		return float64(nums[(l1+l2)/2] + nums[(l1+l2-1)/2]) / 2
	} else {
		return float64(nums[(l1+l2-1)/2])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3}))
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3,4}))
}