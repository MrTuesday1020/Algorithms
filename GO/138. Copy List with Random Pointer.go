package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{
			Val: node.Val,
			Next: node.Next,
		}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	newHead := head.Next
	for node := head; node != nil; node = node.Next {
		newNode := node.Next
		node.Next = node.Next.Next   // 将两个链表拆开 这个是拆出原始链表
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next // 将两个链表拆开 这个是拆出结果链表
		}
		// Question 这里为啥两个拆开动作不能前后叫唤呢？？？
	}
	return newHead
}

func main() {
	
}