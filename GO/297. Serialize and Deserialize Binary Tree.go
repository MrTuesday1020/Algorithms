package main

import "fmt"

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)
	return fmt.Sprintf("%d,%s,%s", root.Val, left, right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var recursion func() *TreeNode
	recursion = func() *TreeNode {
		first := sp[0]
		sp = sp[1:]
		if first == "null" {
			return nil
		}
		val, _ := strconv.Atoi(first)
		// dfs, build left child first, then build right child
		return &TreeNode{
			Val: val,
			Left: recursion(),
			Right: recursion(),
		}
	}
	
	return recursion()
}


/**
* Your Codec object will be instantiated and called as such:
* ser := Constructor();
* deser := Constructor();
* data := ser.serialize(root);
* ans := deser.deserialize(data);
*/

func main() {
	codec := Constructor()
	
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode {
			Val: 2,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 5,
				Left: nil,
				Right: nil,
			},
		},
	}
	
	data := codec.serialize(root)
	fmt.Println(data)
}