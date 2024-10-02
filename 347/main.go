package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	hashMap := make(map[int]int)
	for _, v := range nums {
		if value, found := hashMap[v]; found {
			ctn := value + 1
			hashMap[v] = ctn
		} else {
			hashMap[v] = 1
		}
	}

	arr := [][]int{}
	for key, value := range hashMap {
		arr = append(arr, []int{key, value})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] >= arr[j][1]
	})

	keys := []int{}

	for i := 0; i < k; i++ {
		keys = append(keys, arr[i][0])
	}

	return keys
}

func main() {

	nums1 := []int{1, 1, 1, 2, 2, 3}
	b := topKFrequent(nums1, 2)
	fmt.Println("BVB", b)
	nums2 := []int{1}
	c := topKFrequent(nums2, 1)
	fmt.Println("C", c)
	nums3 := []int{1, 2}
	d := topKFrequent(nums3, 2)
	fmt.Println("D", d)

	nums := []int{4, 1, -1, 2, -1, 2, 3}
	a := topKFrequent(nums, 2)
	fmt.Println("E", a)

	nums5 := []int{5, 3, 1, 1, 1, 3, 73, 1}
	j := topKFrequent(nums5, 2)
	fmt.Println("K", j)
}
