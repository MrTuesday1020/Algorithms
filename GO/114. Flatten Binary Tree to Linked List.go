package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	list := preOrderList(root)
	for i := 0; i < len(list)-1; i++ {
		current := list[i]
		next := list[i+1]
		current.Left = nil
		current.Right = next
	}
}

func preOrderList(root *TreeNode) []*TreeNode {
	var list []*TreeNode
	if root != nil {
		list = append(list, root)
		list = append(list, preOrderList(root.Left)...)
		list = append(list, preOrderList(root.Right)...)
	}
	return list
}

func PrintRet(node *TreeNode){
	if node != nil {
		fmt.Print(node.Val, " ")
		PrintRet(node.Left)
		PrintRet(node.Right)
	} else {
		fmt.Print("null ")
	}
}

func main() {
	n2 := &TreeNode{
			Val: 2,
			Left: nil,
			Right: nil,
		}
	n3 := &TreeNode{
			Val: 3,
			Left: nil,
			Right: nil,
		}
	n1 := &TreeNode{
		Val: 1,
		Left: n2, 
		Right: n3,
	}
	PrintRet(n1)
	fmt.Println()
	fmt.Println("================")
	flatten(n1)
	PrintRet(n1)
	fmt.Println()
}