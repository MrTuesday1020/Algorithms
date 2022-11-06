package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	currentLevel := []*TreeNode{root}
	for len(currentLevel) > 0 {
		currentAns := []int{}
		nextLevel := []*TreeNode{}
		for i := 0; i < len(currentLevel); i++ {
			if currentLevel[i] != nil {
				currentAns = append(currentAns, currentLevel[i].Val)
				if currentLevel[i].Left != nil {
					nextLevel = append(nextLevel, currentLevel[i].Left)
				}
				if currentLevel[i].Right != nil {
					nextLevel = append(nextLevel, currentLevel[i].Right)
				}
			}
		}
		ans = append(ans, currentAns)
		currentLevel = nextLevel
	}
	return ans
}

func main() {
	
}