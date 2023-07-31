package utils_test

import (
	"testing"

	"github.com/Firhan384/restfull-api-golang/app/utils"
	"github.com/stretchr/testify/assert"
)

func TestTextHash(t *testing.T) {
	t.Run("Generate text hash", func(t *testing.T) {
		textHash := "pass"
		hash := utils.TextHash(textHash)

		assert.NotEqual(t, textHash, hash)
	})
}

func TestCompareHash(t *testing.T) {
	t.Run("CompareHash is not compare", func(t *testing.T) {
		hash := utils.TextHash("test")
		text := "tests"
		compare := utils.CompareHash(text, hash)

		assert.Equal(t, false, compare)
	})

	t.Run("CompareHash is compare", func(t *testing.T) {
		hash := utils.TextHash("test")
		text := "test"
		compare := utils.CompareHash(text, hash)

		assert.Equal(t, true, compare)
	})
}
