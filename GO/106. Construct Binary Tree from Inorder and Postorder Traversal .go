package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	length := len(inorder)
	if length == 0 {
		return nil
	}
	i := 0
	for ; i < length; i++ {
		if inorder[i] == postorder[length-1] {
			break
		}
	}
	return &TreeNode{
		Val: postorder[length-1],
		Left: buildTree(inorder[:i], postorder[:i]),
		Right: buildTree(inorder[i+1:], postorder[i:length-1]),
	}
}

func main() {
	
}