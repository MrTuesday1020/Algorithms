package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func rob(root *TreeNode) int {
	var dfs func(*TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			// first value means the ans if node is selected
			// second value means the ans if node is node selected
			return []int{0, 0}
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		// node is selected, children are not selected
		nodeSelected := node.Val + left[1] + right[1]
		// node is not selected = max value of left child + max value of right child
		nodeNotSelected := max(left[0], left[1]) + max(right[0], right[1])
		return []int{nodeSelected, nodeNotSelected}
	}
	ans := dfs(root)
	return max(ans[0], ans[1])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	
}