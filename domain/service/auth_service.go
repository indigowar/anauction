package service

import (
	"context"
	"errors"
	"log/slog"
	"net/mail"

	"github.com/google/uuid"

	"github.com/indigowar/anauction/domain/models"
)

type Auth struct {
	logger *slog.Logger

	userStorage UserStorage
}

func (auth *Auth) SignIn(ctx context.Context, name string, email *mail.Address, password string) (uuid.UUID, error) {
	user, err := models.NewUser(name, email, nil, password)
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := auth.userStorage.Add(ctx, user); err != nil {
		var duplicationErr *DuplicationError
		if errors.As(err, &duplicationErr) {
			return uuid.UUID{}, err
		}

		auth.logger.Error(
			"AuthService.SignIn has FAILED",
			"error", err.Error(),
		)

		return uuid.UUID{}, ErrInternal
	}

	return user.ID(), nil
}

func (auth *Auth) Login(ctx context.Context, email *mail.Address, password string) (uuid.UUID, error) {
	user, err := auth.userStorage.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return uuid.UUID{}, ErrInvalidCredentials
		}

		auth.logger.Error(
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

func NewAuth(logger *slog.Logger, userstorage UserStorage) Auth {
	return Auth{
		logger:      logger,
		userStorage: userstorage,
	}
}
