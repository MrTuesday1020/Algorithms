package main

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	m := make(map[*TreeNode]*TreeNode)    // key child   value parent
	var findParent func(*TreeNode)
	findParent := func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			m[node.Left] = node
			findParent(node.Left)
		}
		if node.Right != nil {
			m[node.Right] = node
			findParent(node.Right)
		}
	}
	findParent(root)
	
	ans := []int{}
	var findAns func(*TreeNode, *TreeNode, int)
	findAns := func(current, last *TreeNode, depth) {
		if current == nil {
			return
		}
		if depth == k {
			ans = append(ans, current.Val)
			return
		}
		if current.Left != last {  // go left
			findAns(current.Left, current, depth + 1)
		}
		if current.Right != last { // go right
			findAns(current.Right, current, depth + 1)
		}
		if m[current] != last { // go up
			findAns(m[current], current, depth + 1)
		}
	}
	findAns(target, nil, 0)
	return ans
}

func main() {
	
}