package main

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
	PrintListNode(head)
}