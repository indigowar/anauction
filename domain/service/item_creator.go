package service

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/indigowar/anauction/domain/models"
)

type ItemCreator struct {
	logger  *slog.Logger
	storage ItemStorage
}

func (im *ItemCreator) CreateItem(
	ctx context.Context,
	owner uuid.UUID,
	name string,
	image *url.URL,
	description string,
	startingPrice float64,
	endTime time.Time,
) (models.Item, error) {
	item, err := models.NewItem(
		owner,
		name,
		image,
		description,
		startingPrice,
		endTime,
	)

	if err != nil {
		return models.Item{}, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}

	if err := im.storage.Add(ctx, item); err != nil {
		// todo: handle the storage error
		return models.Item{}, ErrInternal
	}

	return item, nil
}

func NewItemCreator(logger *slog.Logger, storage ItemStorage) ItemCreator {
	return ItemCreator{
		logger:  logger,
		storage: storage,
	}
}
