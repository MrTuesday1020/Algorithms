package main


func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	ans := rootSum(root, targetSum) // start dfs to find a path from root
	ans += pathSum(root.Left, targetSum) // dfs to left child
	ans += pathSum(root.Right, targetSum) // dfs to right child
	return ans
}

func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	val := root.Val
	ans := 0
	if val == targetSum {
		ans += 1  // 找到一条路径
	}
	// 节点有正有负 需要继续寻找
	ans += rootSum(root.Left, targetSum-val)
	ans += rootSum(root.Right, targetSum-val)
	return ans
}

func main() {
	
}