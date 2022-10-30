package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func PrintListNode(head *ListNode){
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Val: 0,
		Next: nil,
	}
	tmp := head
	
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				tmp.Next = list1
				list1 = list1.Next		
			} else {
				tmp.Next = list2
				list2 = list2.Next
			}
		} else if list1 != nil {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}
		tmp = tmp.Next
	}
	
	return head.Next
}



func main() {
	
}