package main

import "fmt"

// Dado o headvalor de uma lista encadeada e x, particione-a de modo que todos os nós menores que x venham antes dos nós maiores ou iguais a x.

// Você deve preservar a ordem relativa original dos nós em cada uma das duas partições.

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

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, nil}
	current := dummy
	dummy2 := &ListNode{}
	aux := dummy2

	for head != nil {
		if head.Val < x {
			newNode := &ListNode{Val: head.Val}
			current.Next = newNode
			current = current.Next
		} else {
			newNode := &ListNode{Val: head.Val}
			aux.Next = newNode
			aux = aux.Next
		}
		head = head.Next
	}

	row := dummy

	for row != nil && row.Next != nil {
		row = row.Next
	}

	row.Next = dummy2.Next

	return dummy.Next
}

func main() {
	// Entrada: cabeça = [1,4,3,2,5,2], x = 3
	//  Saída: [1,2,2,4,3,5]

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}

	head2 := &ListNode{
		Val:  1,
		Next: nil,
	}

	a := partition(head, 3)
	b := partition(head2, 0)

	for a != nil {
		fmt.Println("aishua", a.Val)
		a = a.Next
	}

	for b != nil {
		fmt.Println("2", b.Val)
		b = b.Next
	}

	// partition(head, 3)
}
