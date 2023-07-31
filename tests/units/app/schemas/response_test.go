package schemas_test

import (
	"testing"

	"github.com/Firhan384/restfull-api-golang/app/schemas"
	"github.com/stretchr/testify/assert"
)

func TestNewResponseApi(t *testing.T) {
	t.Run("If empty", func(t *testing.T) {
		actual := schemas.NewResponseApi()
		expectation := schemas.ResponseApi{
			Error:         false,
			StatusCode:    0,
			Data:          nil,
			DetailMessage: nil,
			Message:       "",
		}

		assert.Equal(t, expectation, actual)
	})

	t.Run("if not empty", func(t *testing.T) {
		actual := schemas.NewResponseApi()
		actual.Error = true

		expectation := schemas.ResponseApi{
			Error: true,
		}

		assert.Equal(t, actual.Error, expectation.Error)
	})
}
