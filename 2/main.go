package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	var arr1 []int
	var arr2 []int
	current1, current2 := dummy1, dummy2

	for l1 != nil || l2 != nil {
		if l1 != nil {
			arr1 = append(arr1, l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			arr2 = append(arr2, l2.Val)
			l2 = l2.Next

		}

	}
	i, j := len(arr1)-1, len(arr2)-1

	for i >= 0 {
		newNode := &ListNode{Val: arr1[i], Next: nil}
		current1.Next = newNode
		current1 = current1.Next
		i--
	}

	for j >= 0 {
		newNode := &ListNode{Val: arr2[j], Next: nil}
		current2.Next = newNode
		current2 = current2.Next
		j--
	}

	current1 = dummy1.Next
	current2 = dummy2.Next
	aux := &ListNode{}
	current := aux
	res := 0
	for current1 != nil || current2 != nil {
		sum := 0 + res
		if current1 != nil {
			sum += current1.Val
			current1 = current1.Next
		}
		if current2 != nil {
			sum += current2.Val
			current2 = current2.Next
		}
		carry := sum / 10
		unit := sum % 10

		newNode := &ListNode{Val: unit, Next: nil}
		current.Next = newNode
		if current1 == nil && current2 == nil && carry > 0 {
			current.Next.Next = &ListNode{Val: carry, Next: nil}
		}

		current = current.Next
		res = carry
	}

	return aux.Next
}

func main() {

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	// l1 := &ListNode{
	// 	Val: 9,
	// 	Next: &ListNode{
	// 		Val: 9,
	// 		Next: &ListNode{
	// 			Val: 9,
	// 			Next: &ListNode{
	// 				Val: 9,
	// 				Next: &ListNode{
	// 					Val: 9,
	// 					Next: &ListNode{
	// 						Val: 9,
	// 						Next: &ListNode{
	// 							Val: 9,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// l2 := &ListNode{
	// 	Val: 9,
	// 	Next: &ListNode{
	// 		Val: 9,
	// 		Next: &ListNode{
	// 			Val: 9,
	// 			Next: &ListNode{
	// 				Val:  9,
	// 				Next: nil,
	// 			},
	// 		},
	// 	},
	// }

	a := addTwoNumbers(l1, l2)

	for a != nil {
		fmt.Println("Asad", a.Val)
		a = a.Next
	}
}
