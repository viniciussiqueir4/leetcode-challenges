package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		str := "PAYPALISHIRING"
		result := convert(str, 3)

		assert.Equal(t, "PAHNAPLSIIGYIR", result)
	})
	t.Run("Test case 2", func(t *testing.T) {
		str := "PAYPALISHIRING"
		result := convert(str, 4)

		assert.Equal(t, "PINALSIGYAHRPI", result)
	})
}
