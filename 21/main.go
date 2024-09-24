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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tail := &ListNode{}
	current := tail
	for list1 != nil || list2 != nil {
		if list1 != nil && (list2 == nil || list1.Val <= list2.Val) {
			current.Next = &ListNode{list1.Val, nil}
			current = current.Next
			list1 = list1.Next
		}
		if list2 != nil && (list1 == nil || list2.Val < list1.Val) {
			current.Next = &ListNode{list2.Val, nil}
			current = current.Next
			list2 = list2.Next
		}
	}

	return tail.Next
}
func main() {
	list1 := &ListNode{Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		}}

	list2 := &ListNode{Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		}}

	l2 := &ListNode{Val: 0}

	rr := mergeTwoLists(nil, l2)

	for rr != nil {
		rr = rr.Next
	}
	fmt.Println("Testcaase1", mergeTwoLists(list1, list2))
}
