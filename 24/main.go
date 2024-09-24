package main

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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	tail := head
	current := tail

	for current != nil {
		if current.Next != nil {
			cp := current.Val
			current.Val = current.Next.Val
			current.Next.Val = cp
			current = current.Next
		}
		current = current.Next
	}

	return tail
}

func main() {
	head := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: nil,
				},
			},
		},
	}

	head2 := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: nil,
			},
		},
	}

	swapPairs(head)
	swapPairs(head2)
	swapPairs(nil)
}
