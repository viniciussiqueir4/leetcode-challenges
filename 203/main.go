package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	tail := &ListNode{0, head}
	current := tail

	for current != nil {
		for current.Next != nil && current.Next.Val == val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}

	return tail.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	head2 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  7,
					Next: nil,
				},
			},
		},
	}

	l1 := removeElements(head, 6)
	el2 := removeElements(head2, 7)

	for l1 != nil {
		fmt.Println("L1", l1.Val)
		l1 = l1.Next
	}

	for el2 != nil {
		fmt.Println("L2", el2.Val)
		el2 = el2.Next
	}

}
