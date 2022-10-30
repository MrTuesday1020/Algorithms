package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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

func isPalindrome(head *ListNode) bool {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for i := 0; i < len(stack) / 2; i++ {
		if stack[i] != stack[len(stack)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	
}