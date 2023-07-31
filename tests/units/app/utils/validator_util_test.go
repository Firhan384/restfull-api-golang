package utils_test

import (
	"testing"

	"github.com/Firhan384/restfull-api-golang/app/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewValidator(t *testing.T) {
	t.Run("NewValidator has validator.Validate", func(t *testing.T) {
		valitate := utils.NewValidator()
		assert.NotNil(t, valitate)
	})
}

func TestValidatorErrors(t *testing.T) {

	type TestStruct struct {
		Name string `validate:"required"`
	}

	t.Run("if Name is empty", func(t *testing.T) {
		testData := TestStruct{
			Name: "",
		}
		validate := utils.NewValidator()
		err := validate.Struct(testData)
		fields := utils.ValidatorErrors(err)

		assert.Contains(t, fields["Name"], "required")
	})
}
