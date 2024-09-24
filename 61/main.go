package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	tail := &ListNode{0, head}
	current := tail
	l := 0
	for head != nil {
		l++
		head = head.Next
	}

	limit := k % l

	for i := 0; i < limit; i++ {

		for current.Next.Next != nil {
			current = current.Next
		}

		tail.Next = &ListNode{current.Next.Val, tail.Next}
		current.Next = nil
		current = tail
	}

	return tail
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
						Val: 5,
					},
				},
			},
		},
	}

	// head2 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val:  3,
	// 			Next: nil,
	// 		},
	// 	},
	// }
	rr := rotateRight(head, 2)
	// rr := rotateRight(head2, 2000000000)

	for rr != nil {
		fmt.Println("Tam", rr.Val)

		rr = rr.Next
	}

}
