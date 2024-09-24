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

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	hashMap := make(map[int]bool)

	for head != nil {
		_, ok := hashMap[head.Val]
		if ok {
			hashMap[head.Val] = true
			if head.Next != nil {
				current.Next = head.Next
			} else {
				current.Next = nil
			}
		} else {
			current.Next = head
			hashMap[head.Val] = false
			current = current.Next
		}

		head = head.Next
	}

	return dummy.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	// 1,1,2,3,3
	head1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}

	head2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	del := deleteDuplicates(head2)
	del2 := deleteDuplicates(head1)
	del3 := deleteDuplicates(head)

	for del != nil {
		fmt.Println("del", del.Val)
		del = del.Next
	}

	for del2 != nil {
		fmt.Println("del2", del2.Val)
		del2 = del2.Next
	}

	for del3 != nil {
		fmt.Println("del3", del3.Val)
		del3 = del3.Next
	}
	// fmt.Println("Hashmap", deleteDuplicates(head2))
}
