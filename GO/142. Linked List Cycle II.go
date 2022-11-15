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

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = slow.Next	// 注意这里要多跑一个节点
	new := head			// 这里用到的性质是 此时head到环点的距离和slow到环点的距离相等
	for new != slow {
		new = new.Next
		slow = slow.Next
	}
	return new
}

func main() {
	
}