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

// remove last n node
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tail := &ListNode{0, head}
	current := tail
	target := n
	for head != nil {
		if target == 1 && head.Next != nil {
			current = current.Next
			target = 2
		}
		target--
  head = head.Next
	}
	current.Next = current.Next.Next
	return tail.Next
}

func main() {
	head := &ListNode{Val: 1,
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
		}}
	// head2 := &ListNode{
	// 	Val:  1,
	// 	Next: nil,
	// }
	// head3 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val:  2,
	// 		Next: nil,
	// 	},
	// }
	// fmt.Println("TestCase2", removeNthFromEnd(head2, 1))
	// fmt.Println("TestCase2", removeNthFromEnd(head3, 2))
	n := removeNthFromEnd(head, 2)

	for n != nil {
		fmt.Println("III", n.Val)
		n = n.Next
	}
}
