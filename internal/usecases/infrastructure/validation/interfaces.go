package validation

import "context"

type Validater interface {
	ValidateUuid(ctx context.Context, target string) error
}
