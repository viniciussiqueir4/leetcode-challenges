package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
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
