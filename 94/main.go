package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var k []int
	if root.Left == nil || (root.Left != nil) {
		k = append(k, inorder(root.Right)...)
	}
	return append(append(inorder(root.Left), root.Val), k...)
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	inorderLeft := append(inorder(root.Left), root.Val)
	inorderRight := inorder(root.Right)
	return append(inorderLeft, inorderRight...)
}

func main() {
	t1 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: nil,
		},
	}

	t2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 9,
				},
			},
		},
	}

	t3 := &TreeNode{
		Val:   5,
		Right: nil,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	x1 := inorderTraversal(t1)
	x2 := inorderTraversal(t2)
	x3 := inorderTraversal(t3)

	fmt.Println("#TESTCASE1", x1)
	fmt.Println("#TESTCASE2", x2)
	fmt.Println("#TESTCASE3", x3)

}
