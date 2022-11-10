package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
	var nodes []*TreeNode
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nodes = append(nodes, node)
		inorder(node.Right)
	}
	inorder(root)  // inorder to build a ordered slice
	
	// find the nodes to swap
	var node1, node2 *TreeNode
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i].Val > nodes[i+1].Val {
			node2 = nodes[i+1]
			if node1 == nil {
				node1 = nodes[i]
			} else {
				break
			}
		}
	}
	
	node1.Val, node2.Val = node2.Val, node1.Val
}

func main() {
	
}