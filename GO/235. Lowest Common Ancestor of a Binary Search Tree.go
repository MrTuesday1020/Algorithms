package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val   int
*     Left  *TreeNode
*     Right *TreeNode
* }
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val {
		return p
	}
	if root.Val == q.Val {
		return q
	}
	if root.Val < p.Val && root.Val < q.Val {
		// current val is smaller than both p and q, then ans is in the right side sub tree
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > p.Val && root.Val > q.Val {
		// current val is greater than both p and q, then ans is in the left side sub tree
		return lowestCommonAncestor(root.Left, p, q)
	}
	// find lowest common ancestor
	return root
}

func main() {
	
}