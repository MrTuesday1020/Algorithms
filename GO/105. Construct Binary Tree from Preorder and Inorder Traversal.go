package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root := &TreeNode{
		Val: preorder[0],
		Left: buildTree(preorder[1:i+1], inorder[:i]),
		Right: buildTree(preorder[i+1:], inorder [i+1:]),
	}
	return root
}

func main() {
	
}