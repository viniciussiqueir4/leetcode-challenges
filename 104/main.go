package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:  15,
				Left: nil,
			},
			Right: &TreeNode{
				Val:  7,
				Left: nil,
			},
		},
	}

	// b := &TreeNode{
	// 	Val:  1,
	// 	Left: nil,
	// 	Right: &TreeNode{
	// 		Val: 2,
	// 	},
	// }

	a := maxDepth(root)
	// c := maxDepth(b)

	fmt.Println("XXX", a)
	// fmt.Println("ssss", c)

}
