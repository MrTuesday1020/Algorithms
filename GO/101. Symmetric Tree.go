package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return recursion(root.Left, root.Right)
}

func recursion(node1 *TreeNode, node2 *TreeNode) {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else if node1.Val != node2.Val {
		return false
	}
	return recursion(node1.Left, node2.Right) && recursion(node1.Right, node2.Left)
}

func main() {
	
}