package order

import (
	"context"
	"errors"
)

var (
	// ErrOrderNotFound represents not found error
	ErrOrderNotFound = errors.New("order not found")

	// ErrCmdRepository represents db conn error
	ErrCmdRepository = errors.New("unable to command repository")

	// ErrQueryRepository represents query error
	ErrQueryRepository = errors.New("unable to query repository")
)

// Service describes the Order service.
type Service interface {
	Create(ctx context.Context, order Order) (string, error)
	GetByID(ctx context.Context, id string) (Order, error)
	ChangeStatus(ctx context.Context, id string, status string) error
}
