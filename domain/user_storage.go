package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound      = errors.New("user is not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserStorage interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	GetByName(ctx context.Context, name string) (User, error)
	Add(ctx context.Context, user User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, user User) error
}
