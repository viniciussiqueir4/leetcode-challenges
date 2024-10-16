package main

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	defeats := make(map[int]int)
	answers := make([][]int, 2)
	for _, v := range matches {
		win := v[0]
		los := v[1]

		if value, exist := defeats[los]; exist {
			defeats[los] = value + 1

			if _, found := defeats[win]; !found {
				defeats[win] = 0
			}
		} else {
			if _, found := defeats[win]; !found {
				defeats[win] = 0
			}
			defeats[los] = 1
		}
	}

	for i, value := range defeats {
		if value == 0 {
			answers[0] = append(answers[0], i)
		} else if value == 1 {
			answers[1] = append(answers[1], i)
		}
	}
	sort.Ints(answers[0])
	sort.Ints(answers[1])

	return answers
}

func main() {
	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	ok := findWinners(matches)
	fmt.Println("#Case1", ok)
}
