package main

import "fmt"

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	left, right := 0, 1

	for left < right {
		if left >= right {
			right++
		}
	}
	return 0
}

func main() {
	fmt.Println("TestCase1", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println("TestCase2", trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println("TestCase3", trap([]int{0}))
	fmt.Println("TestCase4", trap([]int{4, 2, 3}))
}
