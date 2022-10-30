package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func preorderTraversal(root *TreeNode) []int {
	var list []int
	if root != nil {
		list = append(list, root.Val)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}	
	return list
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
	fmt.Println(preorderTraversal(n1))
}