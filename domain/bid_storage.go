package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrBidNotFound      = errors.New("bid is not found")
	ErrBidAlreadyExists = errors.New("bid already exists")
)

type BidStorage interface {
	GetByID(ctx context.Context, id uuid.UUID) (Bid, error)
	GetByUser(ctx context.Context, id uuid.UUID) ([]Bid, error)
	GetByItem(ctx context.Context, id uuid.UUID) ([]Bid, error)
	GetByItemAndUser(ctx context.Context, user uuid.UUID, item uuid.UUID) ([]Bid, error)
	Add(ctx context.Context, bid Bid) error
	Delete(ctx context.Context, id uuid.UUID) error
}
