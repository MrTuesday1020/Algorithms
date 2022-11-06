package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(head1, head2 *ListNode) *ListNode {
	head := &ListNode{}
	temp, temp1, temp2 := head, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return head.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	
	length := 0
	for node := head; node != nil; node = node.Next {
		length += 1
	}
	
	dummy := &ListNode{Next: head}
	for l := 1; l < length; l <<= 1 {
		pre, cur := dummy, dummy.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < l && cur.Next != nil; i++ {
				cur = cur.Next
			}
			
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < l && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			cur = next
			
			pre.Next = merge(head1, head2)
			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	
	return dummy.Next
}

func main() {
	
}