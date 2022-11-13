package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return &TreeNode{
			Val: root2.Val,
			Left: mergeTrees(nil, root2.Left),
			Right: mergeTrees(nil, root2.Right),
		}
	}
	if root2 == nil {
		return &TreeNode{
			Val: root1.Val,
			Left: mergeTrees(root1.Left, nil),
			Right: mergeTrees(root1.Right, nil),
		}
	}
	return &TreeNode{
			Val: root1.Val + root2.Val,
			Left: mergeTrees(root1.Left, root2.Left),
			Right: mergeTrees(root1.Right, root2.Right),
		}
}

func main() {
	
}