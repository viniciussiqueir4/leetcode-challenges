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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	i := 0
	hashMap := make(map[int]int)
	for head != nil {
		i++
		hashMap[i] = head.Val
		head = head.Next
	}

	curr := &ListNode{0, nil}
	newList := curr
	for k := i; k >= 0; k-- {
		if v, found := hashMap[k]; found {
			newList.Next = &ListNode{v, nil}
			newList = newList.Next

		}
	}
	return curr.Next

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

	lst := reverseList(head)

	for lst != nil {
		fmt.Println("JIASa", lst.Val)
		lst = lst.Next
	}
}
