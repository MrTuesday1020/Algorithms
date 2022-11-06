package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	ans := 0
	var recursion func(*TreeNode, int)
	recursion = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth += 1
		ans = max(ans, depth)
		recursion(node.Left, depth)
		recursion(node.Right, depth)
	}
	recursion(root, 0)
	return ans
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	
}