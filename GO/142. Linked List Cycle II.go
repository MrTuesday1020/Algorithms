package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] {
			return head
		} else {
			m[head] = true
			head = head.Next
		}
	}
	return nil
}

func main() {
	
}