package main

// merge from end to start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	m = m - 1
	n = n - 1
	for m >= 0 || n >= 0 {
		if m >= 0 && n >= 0 {
			if nums1[m] > nums2[n] {
				nums1[m+n+1] = nums1[m]
				m = m - 1
			} else {
				nums1[m+n+1] = nums2[n]
				n = n - 1
			}
		} else if m >= 0 {
			nums1[m+n+1] = nums1[m]
			m = m - 1
		} else {
			nums1[m+n+1] = nums2[n]
			n = n - 1
		}
	}
}

func main() {
	
}