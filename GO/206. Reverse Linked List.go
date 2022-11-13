package main

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next    // 拿出下一个节点暂存
		cur.Next = pre      // 将当前节点的下一个节点置为前一个节点(反转)
		pre = cur           // 那么当前节点和前一个节点就组成了一个新的“前一个节点”
		cur = next          // 那么当前节点后移 开始处理下一个节点
	}
	return pre  // 最后“前一个节”点就包含了当前所有的节点 并且已经反转
}

/*
第一轮出栈，head为5，head.next为空，返回5
第二轮出栈，head为4，head.next为5，执行head.next.next=head也就是5.next=4，
			把当前节点的子节点的子节点指向当前节点
			此时链表为1->2->3->4<->5，由于4与5互相指向，所以此处要断开4.next=null
			此时链表为1->2->3->4<-5
			返回节点5
第三轮出栈，head为3，head.next为4，执行head.next.next=head也就是4.next=3，
			此时链表为1->2->3<->4<-5，由于3与4互相指向，所以此处要断开3.next=null
			此时链表为1->2->3<-4<-5
			返回节点5
第四轮出栈，head为2，head.next为3，执行head.next.next=head也就是3.next=2，
			此时链表为1->2<->3<-4<-5，由于2与3互相指向，所以此处要断开2.next=null
			此时链表为1->2<-3<-4<-5
			返回节点5
第五轮出栈，head为1，head.next为2，执行head.next.next=head也就是2.next=1，
			此时链表为1<->2<-3<-4<-5，由于1与2互相指向，所以此处要断开1.next=null
			此时链表为1<-2<-3<-4<-5
			返回节点5
出栈完成，最终头节点5->4->3->2->1
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	new := reverseList(head.Next)
	head.Next.Next = head       // 反转 但是会造成一个环
	head.Next = nil             // 将环断开
	return new
}

func main() {
	
}