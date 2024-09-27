package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left
	return root
}

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  1,
				Left: nil,
			},
			Right: &TreeNode{
				Val:  3,
				Left: nil,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:  6,
				Left: nil,
			},
			Right: &TreeNode{
				Val:  9,
				Left: nil,
			},
		},
	}

	inverted := invertTree(tree)

	fmt.Println("INverterd", inverted.Left.Left, inverted.Right.Left)
}
