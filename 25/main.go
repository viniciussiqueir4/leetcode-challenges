package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) AddNode(val int) {
	head := &ListNode{
		Val:  val,
		Next: nil,
	}

	if l.Next != nil {
		l.Next = head
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	ct := 1
	_, cp := []int{}, []int{}

	// reverseArray := func(arr []int) []int {
	// 	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
	// 		arr[i], arr[j] = arr[j], arr[i]
	// 	}
	// 	return arr
	// }

	for head != nil {
		cp = append(cp, head.Val)

		if ct == k {
			ct = 0
			cp = []int{}
		}
		ct++
		head = head.Next
	}

	return &ListNode{}
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
	reverseKGroup(&list, 2)
}
