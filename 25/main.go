package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() {
	if s == nil {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Clear() {
	if s == nil {
		return
	}
	s.items = []int{}
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{}
	var aux = dummy
	stack := &Stack{}

	for head != nil {
		var dummyAux = &ListNode{}
		var lstAux = dummyAux
		for j := 0; j < k; j++ {
			stack.Push(head.Val)
			dummyAux.Next = &ListNode{Val: head.Val}
			dummyAux = dummyAux.Next
			head = head.Next

			if head == nil && j+1 != k {
				dummy.Next = lstAux.Next
				stack.Clear()
				break
			}
		}

		for !stack.IsEmpty() {
			v, _ := stack.Top()
			dummy.Next = &ListNode{Val: v}
			dummy = dummy.Next
			stack.Pop()
		}

		stack.Clear()
	}

	return aux.Next
}

func main() {
	list := ListNode{Val: 1, Next: &ListNode{
		Val: 2,
	}}
	revert := reverseKGroup(&list, 2)

	for revert != nil {
		fmt.Println("KKK", revert.Val)
		revert = revert.Next
	}
}
