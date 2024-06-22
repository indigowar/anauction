package postgres

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/anauction/domain/models"
	"github.com/indigowar/anauction/domain/service"
	"github.com/indigowar/anauction/storage/postgres/data"
)

type ItemStorage struct {
	queries *data.Queries
}

var _ service.ItemStorage = &ItemStorage{}

// Add implements service.ItemStorage.
func (i *ItemStorage) Add(ctx context.Context, item models.Item) error {
	_, err := i.queries.InsertItem(ctx, itemToInsertItemArg(item))
	if err != nil {
		if err := checkDuplicationError(err); err != nil {
			return err
		}

		if occurred, field := checkForeignKeyViolationError(err); occurred {
			return fmt.Errorf("object of %s not found", field)
		}

		return err
	}
	return nil
}

// Delete implements service.ItemStorage.
func (i *ItemStorage) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := i.queries.DeleteItem(ctx, pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return service.ErrItemNotFound
		}

		return err
	}

	return nil
}

// GetByID implements service.ItemStorage.
func (i *ItemStorage) GetByID(ctx context.Context, id uuid.UUID) (models.Item, error) {
	data, err := i.queries.GetItemByID(ctx, pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Item{}, service.ErrItemNotFound
		}

		return models.Item{}, err
	}

	return dataItemToModel(data), nil
}

// GetByOwner implements service.ItemStorage.
func (i *ItemStorage) GetByOwner(ctx context.Context, id uuid.UUID) ([]models.Item, error) {
	data, err := i.queries.GetItemsByOwner(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, service.ErrItemNotFound
		}

		return nil, err
	}

	result := make([]models.Item, len(data))
	for i, v := range data {
		result[i] = dataItemToModel(v)
	}

	return result, nil
}

// Update implements service.ItemStorage.
func (i *ItemStorage) Update(ctx context.Context, item models.Item) error {
	_, err := i.queries.UpdateItem(ctx, data.UpdateItemParams{
		ID:          pgtype.UUID{Bytes: item.ID(), Valid: true},
		Owner:       pgtype.UUID{Bytes: item.Owner(), Valid: true},
		Name:        item.Name(),
		Image:       item.Image().String(),
		Description: item.Description(),
		StartPrice:  item.StartingPrice(),
		CreatedAt: pgtype.Timestamp{
			Time:             item.CreatedAt(),
			InfinityModifier: 0,
			Valid:            true,
		},
		ClosedAt: pgtype.Timestamp{
			Time:             item.ClosedAt(),
			InfinityModifier: 0,
			Valid:            true,
		},
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return service.ErrItemNotFound
		}

		if err := checkDuplicationError(err); err != nil {
			return err
		}

		if occurred, field := checkForeignKeyViolationError(err); occurred {
			return fmt.Errorf("object of %s not found", field)
		}

		return err
	}

	return nil
}

func NewItemStorage(conn *pgx.Conn) *ItemStorage {
	return &ItemStorage{
		queries: data.New(conn),
	}
}

func itemToInsertItemArg(i models.Item) data.InsertItemParams {
	return data.InsertItemParams{
		ID:          pgtype.UUID{Bytes: i.ID(), Valid: true},
		Owner:       pgtype.UUID{Bytes: i.Owner(), Valid: true},
		Name:        i.Name(),
		Image:       i.Image().String(),
		Description: i.Description(),
		StartPrice:  i.StartingPrice(),
		CreatedAt: pgtype.Timestamp{
			Time:             i.CreatedAt(),
			InfinityModifier: 0,
			Valid:            true,
		},
		ClosedAt: pgtype.Timestamp{
			Time:             i.ClosedAt(),
			InfinityModifier: 0,
			Valid:            true,
		},
	}
}

func dataItemToModel(data data.Item) models.Item {
	image, _ := url.Parse(data.Image)

	return models.NewRawItem(
		data.ID.Bytes,
		data.Owner.Bytes,
		data.Name,
		image,
		data.Description,
		data.StartPrice,
		data.CreatedAt.Time,
		data.ClosedAt.Time,
	)
}
