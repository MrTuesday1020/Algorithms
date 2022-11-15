package main

func removeElements(head *ListNode, val int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	for node := pre; node.Next != nil; {
		if node.Next.Val == val {
			node.Next = node.Next.Next  // if the val of next node is equl to val, jump over it
		} else {
			node = node.Next
		}
	}
	return pre.Next
}

func main() {
	
}