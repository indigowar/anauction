package service

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/indigowar/anauction/domain/models"
)

var (
	ErrItemNotFound = errors.New("item was not found")
)

type ItemStorage interface {
	GetByID(ctx context.Context, id uuid.UUID) (models.Item, error)
	GetByOwner(ctx context.Context, id uuid.UUID) ([]models.Item, error)
	Add(ctx context.Context, item models.Item) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, item models.Item) error
}
