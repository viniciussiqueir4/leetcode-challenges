package main

import "fmt"

type Three struct {
	Head  int
	Left  *Three
	Right *Three
}

func (t *Three) insert(val int) {
	if val <= t.Head {
		if t.Left == nil {
			t.Left = &Three{Head: val}
		} else {
			t.Left.insert(val)
		}
	} else {
		if t.Right == nil {
			t.Right = &Three{Head: val}
		} else {
			t.Right.insert(val)
		}
	}
}

func (t *Three) Contains(val int) bool {
	if t == nil {
		return false
	}

	if val == t.Head {
		return true
	} else if val < t.Head {
		return t.Left.Contains(val)
	} else {
		return t.Right.Contains(val)
	}
}

func main() {
	root := &Three{Head: 5}
	root.insert(3)
	root.insert(2)
	root.insert(4)

	exist := root.Contains(2)
	fmt.Println("Existe?", exist)
}
