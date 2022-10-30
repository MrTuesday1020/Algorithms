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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var stack []*ListNode
	tmp := &ListNode{0, head}
	for node := tmp; node != nil; node = node.Next {
		stack = append(stack, node)
	}
	pre := stack[len(stack)-1-n] 	// find previous one
	pre.Next = pre.Next.Next    	// delete
	return tmp.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}
	PrintListNode(removeNthFromEnd(head, 2))
}