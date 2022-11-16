package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	l := dummy
	for i := 1; i < left; i++ {    // 前缀个数是left-1
		l = l.Next
	}
	
	r := l.Next
	var pre *ListNode
	cur := r
	for i := left; i <= right; i++ {  // 需要翻转的节点个数是 right-left+1
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	
	l.Next = pre    // 将前半部分和反转部分串联 pre是前半部分不需要排序的部分
	r.Next = cur    // 将后半部分和反转部分串联 cur是后半部分不需要排序的开头了
	return dummy.Next
}

func main() {
	
}