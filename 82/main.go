package main

import "fmt"

// 82. Remove Duplicates from Sorted List II
// Medium
// Topics
// Companies
// Given the head of a sorted linked list, delete all nodes
// that have duplicate numbers, leaving only distinct numbers from the original list.
//  Return the linked list sorted as well.

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

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	current := dummy
	hashMap := make(map[int]int)
	for head != nil {
		value, exist := hashMap[head.Val]
		if exist {
			hashMap[head.Val] = value + 1
		} else {
			hashMap[head.Val] = 1
			if head.Next != nil && head.Val != head.Next.Val {
				current.Next = &ListNode{Val: head.Val, Next: nil}
				current = current.Next
			} else if head.Next == nil {
				current.Next = &ListNode{Val: head.Val, Next: nil}
			}
		}
		head = head.Next
	}

	return dummy.Next
}

func main() {
	// head1 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 3,
	// 				Next: &ListNode{
	// 					Val: 4,
	// 					Next: &ListNode{
	// 						Val: 4,
	// 						Next: &ListNode{
	// 							Val:  5,
	// 							Next: nil,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// head2 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 1,
	// 		Next: &ListNode{
	// 			Val: 1,
	// 			Next: &ListNode{
	// 				Val: 2,
	// 				Next: &ListNode{
	// 					Val:  3,
	// 					Next: nil,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	head3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	// deleteDuplicates(head1)
	// deleteDuplicates(head2)
	a := deleteDuplicates(head3)

	for a != nil {
		fmt.Println("ashuda", a.Val)
		a = a.Next
	}

}
