package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func hasCycle(head *ListNode) bool {
	// using map
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] {
			return true
		} else {
			m[head] = true
			head = head.Next
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	
}