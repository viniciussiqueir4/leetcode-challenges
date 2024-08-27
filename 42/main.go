package main

import "fmt"

func trap(height []int) int {
	left, right := 0, 1
	lim := len(height)
	res := 0
	for right < lim {
		// 1 > 0
		if height[left] > height[right] {
			left = right // 1
		}

		right++ // 2

		// 2 + 1 + 0 + 1
		// else {
		// 	if right+1 == lim {
		// 		left++
		// 		right = left + 1
		// 	} else {
		// 		right++
		// 	}
		// }
	}

	fmt.Println("Hei", res)

	return res
}

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	// trap([]int{4, 2, 0, 3, 2, 5})
}
