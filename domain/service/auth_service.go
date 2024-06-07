package service

import (
	"context"
	"errors"
	"log/slog"
	"net/mail"

	"github.com/google/uuid"

	"github.com/indigowar/anauction/domain/models"
)

type AuthService struct {
	logger *slog.Logger

	userStorage UserStorage
}

func (svc *AuthService) SignIn(ctx context.Context, name string, email *mail.Address, password string) (uuid.UUID, error) {
	user, err := models.NewUser(name, email, nil, password)
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := svc.userStorage.Add(ctx, user); err != nil {
		var duplicationErr *DuplicationError
		if errors.As(err, &duplicationErr) {
			return uuid.UUID{}, err
		}

		svc.logger.Error(
			"AuthService.SignIn has FAILED",
			"error", err.Error(),
		)

		return uuid.UUID{}, ErrInternal
	}

	return user.ID(), nil
}

func (svc *AuthService) Login(ctx context.Context, email *mail.Address, password string) (uuid.UUID, error) {
	user, err := svc.userStorage.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return uuid.UUID{}, ErrInvalidCredentials
		}

		svc.logger.Error(
			"AuthService.Login has FAILED",
			"error", err.Error(),
		)

		return uuid.UUID{}, ErrInternal
	}

	if !user.ComparePassword(password) {
		return uuid.UUID{}, ErrInvalidCredentials
	}

	return user.ID(), nil
}

func NewAuthService(logger *slog.Logger, userstorage UserStorage) AuthService {
	return AuthService{
		logger:      logger,
		userStorage: userstorage,
	}
}
