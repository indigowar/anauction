package service

import (
	"context"
	"errors"
	"net/mail"

	"github.com/google/uuid"

	"github.com/indigowar/anauction/domain/models"
)

var (
	ErrUserNotFound = errors.New("user is not found")
)

type UserStorage interface {
	GetByID(ctx context.Context, id uuid.UUID) (models.User, error)
	GetByEmail(ctx context.Context, email *mail.Address) (models.User, error)
	Add(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, user models.User) error
}
