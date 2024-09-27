package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	tail := head
	current := tail
	for current != nil {
		// if head.Next != nil && head.Next.Val > head.Val {
		// 	prev =
		// }

		current = current.Next
	}
	fmt.Println("DUumy", prev)
	return head
}

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	sortList(head)
}
