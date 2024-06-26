package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"log/slog"
	"math/big"
	"time"

	"github.com/google/uuid"

	"github.com/indigowar/anauction/domain/models"
)

type ItemCreator struct {
	logger       *slog.Logger
	storage      ItemStorage
	imageStorage ImageStorage
}

func (im *ItemCreator) CreateItem(
	ctx context.Context,
	owner uuid.UUID,
	name string,
	image []byte,
	description string,
	startingPrice float64,
	endTime time.Time,
) (models.Item, error) {
	imageKey := im.generateImageKey()
	if err := im.imageStorage.Add(ctx, imageKey, image); err != nil {
		return models.Item{}, ErrInternal
	}

	item, err := models.NewItem(
		owner,
		name,
		imageKey,
		description,
		startingPrice,
		endTime,
	)

	if err != nil {
		return models.Item{}, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}

	if err := im.storage.Add(ctx, item); err != nil {
		return models.Item{}, ErrInternal
	}

	return item, nil
}

func (im *ItemCreator) generateImageKey() string {
	const length = 32
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		result[i] = charset[randomInt.Int64()]
	}
	return string(result)
}

func NewItemCreator(
	logger *slog.Logger,
	storage ItemStorage,
	imageStorage ImageStorage,
) ItemCreator {
	return ItemCreator{
		logger:       logger,
		storage:      storage,
		imageStorage: imageStorage,
	}
}
