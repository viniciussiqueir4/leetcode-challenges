package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) addNodeTree(val int) {
	if val <= t.Val {
		if t.Left != nil {
			t.Left.addNodeTree(val)
		} else {
			t.Left = &TreeNode{Val: val}
		}
	} else {
		if t.Right != nil {
			t.Right.addNodeTree(val)
		} else {
			t.Right = &TreeNode{Val: val}
		}
	}
}

func sortedListToBST(head *ListNode) *TreeNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	h := len(arr) / 2
	newTree := &TreeNode{Val: arr[h]}
	leftRight := append(arr[:h], arr[h+1:]...)

	// func(val int) {
	// 	return
	// }
	// add := func(val int) {
	// 	if val <= newTree.Val {
	// 		if newTree.Left != nil {

	// 		} else {
	// 			t.Left = &TreeNode{Val: val}
	// 		}
	// 	} else {
	// 		if t.Right != nil {
	// 			t.Right.addNodeTree(val)
	// 		} else {
	// 			t.Right = &TreeNode{Val: val}
	// 		}
	// 	}
	// }
	fmt.Println("SDAS", len(leftRight))
	for i := len(leftRight) - 1; i >= 0; i-- {
		newTree.addNodeTree(leftRight[i])
	}

	return newTree
}

func main() {
	head := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	}

	sortedListToBST(head)
}
