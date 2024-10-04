package main

import (
	"fmt"
	"sort"
)

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	left, right := 0, len(skill)-1
	target := skill[0] + skill[right]
	var teams [][]int
	for left < right {
		if target != skill[left]+skill[right] {
			return -1
		}
		teams = append(teams, []int{skill[left], skill[right]})
		left++
		right--
	}

	sum := 0
	for _, value := range teams {
		sum += value[0] * value[1]
	}

	return int64(sum)
}

func main() {
	k := dividePlayers([]int{3, 2, 5, 1, 3, 4})
	q := dividePlayers([]int{3, 4})

	fmt.Println("Testacase1", k)
	fmt.Println("Testacase2", q)

}
