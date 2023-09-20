package gotesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {
	t.Run("negative test case", func(t *testing.T) {
		c := Absolute(-5)
		assert.Equal(t, 5, c, "expect 5, got %d", c)
	})

	t.Run("postive test case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c)
	})

}

func TestAdd(t *testing.T) {
	t.Run("negative add negative", func(t *testing.T) {
		c := Add(-1, -3)
		assert.Equal(t, -4, c)
	})

	t.Run("postive add negative", func(t *testing.T) {
		c := Add(3, -4)
		assert.Equal(t, -1, c)
	})

}

func TestAdd_WithTestTable(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "negative and negative",
			a:        -1,
			b:        -1,
			expected: -2,
		},
		{
			name:     "positive and negative",
			a:        2,
			b:        -1,
			expected: 1,
		},
		{
			name:     "positive and positive",
			a:        3,
			b:        5,
			expected: 8,
		},
		{
			name:     "negative and positive",
			a:        -5,
			b:        2,
			expected: -3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, c)
		})
	}

}
