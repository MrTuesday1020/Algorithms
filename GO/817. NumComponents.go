package main

import (
	"fmt"
)

//You are given the head of a linked list containing unique integer values and 
//an integer array nums that is a subset of the linked list values. 
//
// Return the number of connected components in nums where two values are 
//connected if they appear consecutively in the linked list. 
//
// 
// Example 1: 
// 
// 
//Input: head = [0,1,2,3], nums = [0,1,3]
//Output: 2
//Explanation: 0 and 1 are connected, so [0, 1] and [3] are the two connected 
//components.
// 
//
// Example 2: 
// 
// 
//Input: head = [0,1,2,3,4], nums = [0,3,1,4]
//Output: 2
//Explanation: 0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4
//] are the two connected components.
// 
//
// 
// Constraints: 
//
// 
// The number of nodes in the linked list is n. 
// 1 <= n <= 10⁴ 
// 0 <= Node.val < n 
// All the values Node.val are unique. 
// 1 <= nums.length <= n 
// 0 <= nums[i] < n 
// All the values of nums are unique. 
// 
//
// Related Topics 数组 哈希表 链表 👍 160 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}


func numComponents(head *ListNode, nums []int) int {
	m := make(map[int]bool, len(nums))
	for i := range nums {
		m[nums[i]] = true
	}
	
	var count int = 0
	var flag bool = false
	for head != nil {
		if m[head.Val] && flag == false {
			// find one
			flag = true
			count += 1
		} 
//		if m[head.Val] && flag == true {
//			// already got, do nothing
//		}
		if !m[head.Val] {
			// break
			flag = false
		}
		head = head.Next
	}
	return count
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {
	n5 := &ListNode{Val: 5, Next: nil}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	nums := []int{1,2,4,5}
	fmt.Println(numComponents(n1, nums))
}