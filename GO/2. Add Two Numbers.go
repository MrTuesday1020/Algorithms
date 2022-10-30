package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import(
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {	
	return recursion(l1, l2, 0)
}

func recursion(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	
	var num1, num2 int
	if l1 != nil {
		num1 = l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		num2 = l2.Val
		l2 = l2.Next
	}
	fmt.Println(num1, num2, carry)
	current := (num1 + num2 + carry) % 10
	carry = (num1 + num2 + carry) / 10
	
	return &ListNode{
		Val: current,
		Next: recursion(l1, l2, carry),
	}
}

func PrintRet(l *ListNode){
	for l != nil {
		fmt.Print(l.Val, " ")
		l = l.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: nil,
			},
		},
	}
	PrintRet(addTwoNumbers(l1, l2))
}