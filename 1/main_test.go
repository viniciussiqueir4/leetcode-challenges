package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		res := twoSum([]int{2, 7, 11, 15}, 9)
		assert.Equal(t, []int{0, 1}, res)
	})
	t.Run("Test Case 2", func(t *testing.T) {
		res := twoSum([]int{3, 2, 4}, 6)
		assert.Equal(t, []int{1, 2}, res)
	})
	t.Run("Test Case 3", func(t *testing.T) {
		res := twoSum([]int{3, 3}, 6)
		assert.Equal(t, []int{0, 1}, res)
	})
}
