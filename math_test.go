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
