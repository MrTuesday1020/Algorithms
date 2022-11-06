package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var ans []int

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	recursion(root)
	return ans
}

func recursion(node *TreeNode) {
	if node != nil {
		ans = append(ans, node.Val)
		recursion(node.Left)
		recursion(node.Right)
	}
}

func main() {
	
}