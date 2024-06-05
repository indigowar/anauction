package domain

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type AuthService struct {
	userStorage UserStorage

	logger *slog.Logger
}

func (svc *AuthService) SignIn(ctx context.Context, name string, password string) (uuid.UUID, error) {
	user, err := NewUser(name, password)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%w: %w", ErrInvalidData, err)
	}

	if err := svc.userStorage.Add(ctx, user); err != nil {
		if errors.Is(err, ErrBidAlreadyExists) {
			return uuid.UUID{}, ErrUsernameIsTaken
		}
		logInternalError(svc.logger, "AuthService", "SignIn", "UserStorage", "Add", err)
		return uuid.UUID{}, err
	}

	return user.ID(), nil
}

func (svc *AuthService) Login(ctx context.Context, name string, password string) (uuid.UUID, error) {
	user, err := svc.userStorage.GetByName(ctx, name)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return uuid.UUID{}, ErrInvalidData
		}
		logInternalError(svc.logger, "AuthService", "Login", "UserStorage", "GetByName", err)
		return uuid.UUID{}, err
	}

	if !user.ComparePassword(password) {
		return uuid.UUID{}, ErrInvalidData
	}

	return user.ID(), nil
}

func NewAuthService(userStorage UserStorage, logger *slog.Logger) *AuthService {
	return &AuthService{
		userStorage: userStorage,
		logger:      logger,
	}
}
