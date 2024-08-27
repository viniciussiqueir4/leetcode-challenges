package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) AddNode(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}

	if l == nil {
		l = newNode
		return
	}
	fmt.Println("Oi2", l)

	current := l
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var aux *ListNode
	// var aux2 *ListNode
	count := 1
	aux.AddNode(2)
	fmt.Println("LL", aux)
	fmt.Println("LL", &head)

	for head != nil {
		// if k == count || head.Next != nil {
		// 	for aux != nil {
		// 		aux2.AddNode(aux.Val)
		// 	}
		// 	count = 0
		// 	aux = &ListNode{}
		// }
		fmt.Println("Add Value", head.Val)
		// aux = aux.AddNode(1)
		count++
		head = head.Next
	}

	fmt.Println("Aux", aux)
	// fmt.Println("FNn", *aux)

	return head
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
	fmt.Println("Result", reverseKGroup(&list, 2))
}
