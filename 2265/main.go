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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return count(root.Left) + count(root.Right) + 1
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return root.Val + sum(root.Left) + sum(root.Right)
}

func averageOfSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	cnt := count(root)
	sum := sum(root)
	equal := 0
	if root.Val == sum/cnt {
		equal = 1
	}

	return equal + averageOfSubtree(root.Left) + averageOfSubtree(root.Right)
}

func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	k := averageOfSubtree(tree)
	fmt.Println("KK", k)
}
