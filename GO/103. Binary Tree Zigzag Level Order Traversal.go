package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	level := []*TreeNode{root}
	depth := 0
	ans := [][]int{}
	for len(level) > 0 {
		current := []int{}
		newLevel := []*TreeNode{}
		for _, node := range level {
			if node == nil {
				continue
			}
			if depth % 2 == 0 {
				current = append(current, node.Val)   // append
			} else {
				current = append([]int{node.Val}, current...)  // prepend
			}
			newLevel = append(newLevel, node.Left)
			newLevel = append(newLevel, node.Right)
		}
		depth += 1
		if len(current) > 0 {
			ans = append(ans, current)
		}
		level = newLevel
	}
	return ans
}

func main() {
	
}