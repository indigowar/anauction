package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrItemNotFound      = errors.New("item is not found")
	ErrItemAlreadyExists = errors.New("item already exists")
)

type ItemStorage interface {
	GetByID(ctx context.Context, id uuid.UUID) (Item, error)
	GetByOwner(ctx context.Context, owner uuid.UUID) ([]Item, error)
	Add(ctx context.Context, item Item) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, item Item) error
}
