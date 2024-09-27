package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func split(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && split(left.Left, right.Right) && split(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return split(root.Left, root.Right)
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  3,
				Left: nil,
			},
			Right: &TreeNode{
				Val:  4,
				Left: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  4,
				Left: nil,
			},
			Right: &TreeNode{
				Val:  3,
				Left: nil,
			},
		},
	}

	tree2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:  3,
				Left: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:  3,
				Left: nil,
			},
		},
	}

	v := isSymmetric(tree)
	v2 := isSymmetric(tree2)

	fmt.Println("VV", v)
	fmt.Println("VV2", v2)

}
