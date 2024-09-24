package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	tail := &ListNode{0, head}
	current := tail
	hashMap := make(map[int]int)
	i, j := 1, 1
	for head != nil {
		if i >= left && i <= right {
			hashMap[right+1-j] = head.Val
			j++
		}
		head = head.Next
		i++
	}

	i = 1
	for current != nil {
		if v, found := hashMap[i]; found {
			current.Next.Val = v
		}
		current = current.Next
		i++
	}

	return tail.Next
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
				},
			},
		},
	}

	reverseBetween(head, 1, 4)

}
