package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	smallHead, largeHead := small, large
	for node := head; node != nil; node = node.Next {
		if node.Val < x {
			small.Next = node
			small = small.Next
		} else {
			large.Next = node
			large = large.Next
		}
	}
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}

func main() {
	
}