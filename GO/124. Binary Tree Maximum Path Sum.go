package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt64
	var recursion func(*TreeNode) int
	recursion = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		
		leftSum := max(recursion(node.Left), 0)
		rightSum := max(recursion(node.Right), 0)
		
		currentSum := node.Val + leftSum + rightSum  // sum include root
		
		ans = max(ans, currentSum)
		
		return node.Val + max(leftSum, rightSum)     // sum only include one child
	}
	recursion(root)
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