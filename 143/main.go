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

func reorderList(head *ListNode) {
	tail := head
	current := tail

	// dummy := &ListNode{0, nil}
	// currentDummy := dummy
	_, j := 0, 0

	for current.Next.Next != nil {
		current = current.Next
		j++
	}

	fmt.Println("AG", current)

	for head != nil {
		fmt.Println("asduad", current.Val)
		// current.Next = current

		// fmt.Println("AG", current)
		head = head.Next
	}

	// for head != nil {
	// 	if current.Next != nil {
	// for current.Next.Next != nil {
	// 	current = current.Next
	// 	j++
	// }

	// 		var nextNode *ListNode
	// 		if (k == 0 && j == 0) || k != j {
	// 			nextNode = &ListNode{Val: current.Next.Val}
	// 		}

	// 		currentDummy.Next = &ListNode{Val: head.Val, Next: nextNode}
	// 		currentDummy = currentDummy.Next.Next
	// 		current.Next = current
	// 		current = tail
	// 		j = 1
	// 		k++
	// 	}
	// 	head = head.Next
	// }

	// if dummy.Next != nil {
	// 	current.Next = dummy.Next.Next
	// }
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
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
		},
	}

	// head2 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val:  2,
	// 		Next: nil,
	// 	},
	// }

	reorderList(head)
	// reorderList(head2)

}
