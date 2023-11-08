package input

import (
	"context"
)

type DeleteProductUseCasePort interface {
	Inactive(ctx context.Context, number int) error
}
