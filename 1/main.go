package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	i := 0
	var sums []int
	for len(sums) != 2 {
		value, exist := hashMap[nums[i]]
		if exist {
			sums = append(sums, value, i)
		}
		hashMap[target-nums[i]] = i
		i++
	}
	return sums
}

func main() {
	fmt.Println("Hello", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("Hello", twoSum([]int{3, 2, 4}, 6))
	fmt.Println("Hello", twoSum([]int{3, 3}, 6))

}
