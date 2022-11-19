package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var postorder func(*TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		postorder(root.Left)
		postorder(root.Right)
		ans = append(ans, root.Val)
	}
	postorder(root)
	return ans
}

func main() {
	
}