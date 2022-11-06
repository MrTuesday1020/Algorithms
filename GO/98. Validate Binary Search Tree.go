package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return recursion(root, math.MinInt64, math.MaxInt64)
}

func recursion(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if  node.Val <= lower || node.Val >= upper {
		return false
	}
	return recursion(node.Left, lower, node.Val) && recursion(node.Right, node.Val, upper)
}

func main() {
	
}