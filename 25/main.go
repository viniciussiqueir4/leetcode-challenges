package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	aux := &ListNode{0, head}
	current := aux
	// i := 0

	for head != nil {
		for j := 0; j < k; j++ {
			current = current.Next
		}
		head = head.Next
	}
	return nil
}

func main() {
	list := ListNode{Val: 1, Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}}
	fmt.Println("Result", reverseKGroup(&list, 2))
}
