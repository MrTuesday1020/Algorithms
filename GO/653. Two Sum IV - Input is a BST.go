package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		target :=  k - node.Val
		if m[target] {
			return true
		}
		m[node.Val] = true
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}

func main() {
	
}