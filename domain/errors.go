package domain

import (
	"errors"
	"log/slog"
)

var (
	ErrInvalidData     = errors.New("invalid data was provided")
	ErrUsernameIsTaken = errors.New("provided username is already in use")
)

func logInternalError(
	logger *slog.Logger,
	service string, // Name of the service where the error occurred
	action string, // Action of the service where the error occurred
	source string, // A Source of the error(place where it occurred
	sourceAction string, // Source's action that returned an error
	err error,
) {
	logger.Warn(
		"Internal Server error",
		"service", service,
		"action", action,
		"source", source,
		"source-action", sourceAction,
		"error", err,
	)
}
