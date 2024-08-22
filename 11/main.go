package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left != right {
		width := right - left
		leftIsSmaller, h := height[left] < height[right], 0

		if leftIsSmaller {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		area := h * width
		if area > max {
			max = area
		}
	}
	return max
}

func main() {
	fmt.Println("#TestCase1", maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println("TestCase2", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("TestCase3", maxArea([]int{1, 1}))
}
