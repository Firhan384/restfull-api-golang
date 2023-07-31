package utils_test

import (
	"testing"

	"github.com/Firhan384/restfull-api-golang/app/schemas"
	"github.com/Firhan384/restfull-api-golang/app/utils"
	"github.com/stretchr/testify/assert"
)

func TestStatusText(t *testing.T) {
	t.Run("if status code not found", func(t *testing.T) {
		actual := utils.StatusText(600)

		assert.Equal(t, "", actual)
	})

	t.Run("if status code found", func(t *testing.T) {
		actual := utils.StatusText(schemas.StatusOK)
		expectation := schemas.StatusTextMap[schemas.StatusOK]

		assert.Equal(t, expectation, actual)
	})
}
