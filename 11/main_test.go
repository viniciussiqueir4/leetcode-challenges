package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		heights := []int{4, 3, 2, 1, 4}
		assert.Equal(t, int(16), maxArea(heights))
	})
	t.Run("Test Case 2", func(t *testing.T) {
		heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		assert.Equal(t, int(49), maxArea(heights))
	})
	t.Run("Test Case 3", func(t *testing.T) {
		heights := []int{1, 1}
		assert.Equal(t, int(1), maxArea(heights))
	})
}
