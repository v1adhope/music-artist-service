package validation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v1adhope/music-artist-service/internal/usecases/infrastructure/validation"
)

func prepare() validation.Validater {
	return validation.New()
}

func TestValidateUuidPositive(t *testing.T) {
	validator := prepare()

	tcs := []struct {
		key   string
		input string
	}{
		{
			key:   "Case 1",
			input: "1e712685-714f-6720-a23a-c90103f70be6",
		},
		{
			key:   "Case 2",
			input: "1e71269a-b116-6740-a694-68c004266291",
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			ok := validator.IsValidUuid(tc.input)

			assert.True(t, ok, tc.key)
		})
	}
}

func TestValidateUuidNegative(t *testing.T) {
	validate := prepare()

	tcs := []struct {
		key   string
		input string
	}{
		{
			key:   "Case 1",
			input: "1",
		},
		{
			key:   "Case 2",
			input: "1e71269a-b116-6740-a694",
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			ok := validate.IsValidUuid(tc.input)

			assert.False(t, ok, tc.key)
		})
	}
}
