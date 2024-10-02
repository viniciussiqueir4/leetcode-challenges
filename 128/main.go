package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	l := len(nums) - 1
	if l < 1 {
		return len(nums)
	}

	sort.Ints(nums)
	var arr []int
	cp := []int{}
	for i, v := range nums {
		if i != l {
			if v+1 == nums[i+1] || v == nums[i+1] {
				if v+1 == nums[i+1] {
					arr = append(arr, v)
				}
			} else {
				// fmt.Println("ass", len)
				if len(cp) < len(arr) {
					cp = arr
					arr = []int{}
				} else {
					// cp = arr
					arr = []int{}
					// fmt.Println("Val", cp, arr)

				}
			}
		}
	}
	fmt.Println("aida", cp, arr)
	if len(cp) > len(arr) {
		return len(cp) + 1
	}
	return len(arr) + 1
}

func main() {
	nms := []int{100, 4, 200, 1, 3, 2}
	a := longestConsecutive(nms)
	fmt.Println("K", a)
	// nms2 := []int{7, -9, 3, -6, 3, 5, 3, 6, -2, -5, 8, 6, -4, -6, -4, -4, 5, -9, 2, 7, 0, 0}
	// b := longestConsecutive(nms2)
	// fmt.Println("K1", b)
	// n3 := []int{0, 0}

	// c := longestConsecutive(n3)
	// fmt.Println("K12", c)
}
