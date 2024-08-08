package validation

import (
	"github.com/go-playground/validator/v10"
)

// INFO: Copy from usecases
type Validater interface {
	IsValidUuid(target string) bool
}

var _ Validater = (*Validator)(nil)

type Validator struct {
	*validator.Validate
}

func New() *Validator {
	return &Validator{
		validator.New(validator.WithRequiredStructEnabled()),
	}
}

type uuid struct {
	Value string `validate:"uuid"`
}

func (v *Validator) IsValidUuid(target string) bool {
	uuid := uuid{target}

	if err := v.Struct(uuid); err != nil {
		return false
	}

	return true
}
