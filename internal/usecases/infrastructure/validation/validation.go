package validation

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
)

var _ Validater = (*Validator)(nil)

var (
	ErrNotValidUuid = errors.New("Not valid uuid")
)

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

func (v *Validator) ValidateUuid(ctx context.Context, target string) error {
	uuid := uuid{target}

	if err := v.StructCtx(ctx, uuid); err != nil {
		return ErrNotValidUuid
	}

	return nil
}
