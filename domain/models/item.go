package models

import (
	"errors"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	id    uuid.UUID
	owner uuid.UUID

	name        string
	image       *url.URL
	description string

	startingPrice float64

	createdAt time.Time
	closedAt  time.Time
}

func (i *Item) ID() uuid.UUID          { return i.id }
func (i *Item) Owner() uuid.UUID       { return i.owner }
func (i *Item) Name() string           { return i.name }
func (i *Item) Image() *url.URL        { return i.image }
func (i *Item) Description() string    { return i.description }
func (i *Item) StartingPrice() float64 { return i.startingPrice }
func (i *Item) CreatedAt() time.Time   { return i.createdAt }
func (i *Item) ClosedAt() time.Time    { return i.closedAt }

func (i *Item) IsClosed() bool {
	return time.Now().After(i.closedAt)
}

func (i *Item) SetName(value string) error {
	if len(value) < 3 {
		return errors.New("name should be at least 3 characters long")
	}
	i.name = value
	return nil
}

func (i *Item) SetImage(value *url.URL) error {
	if value == nil {
		return errors.New("image url is not specified")
	}
	i.image = value
	return nil
}

func (i *Item) SetDescription(value string) error {
	i.description = value
	return nil
}

func (i *Item) SetClosedAt(value time.Time) error {
	if time.Now().After(value) {
		return errors.New("closed at time should not be expired")
	}
	i.closedAt = value
	return nil
}

func NewItem(
	owner uuid.UUID,
	name string,
	image *url.URL,
	description string,
	startPrice float64,
	closedAt time.Time,
) (Item, error) {
	if startPrice <= 0 {
		return Item{}, errors.New("starting price has to be a positive number")
	}

	item := Item{
		id:            uuid.New(),
		owner:         owner,
		startingPrice: startPrice,
		createdAt:     time.Now(),
	}

	if err := item.SetName(name); err != nil {
		return Item{}, err
	}

	if err := item.SetImage(image); err != nil {
		return Item{}, err
	}

	if err := item.SetDescription(description); err != nil {
		return Item{}, err
	}

	if err := item.SetClosedAt(closedAt); err != nil {
		return Item{}, err
	}

	return item, nil
}

// NewRawItem creates an item model from provided data without any validation.
//
// This method should not be used in bussiness logic, it is strictly for storage layer
func NewRawItem(
	id uuid.UUID,
	owner uuid.UUID,
	name string,
	image *url.URL,
	description string,
	startingPrice float64,
	createdAt time.Time,
	closedAt time.Time,
) Item {
	return Item{
		id:            id,
		owner:         owner,
		name:          name,
		image:         image,
		description:   description,
		startingPrice: startingPrice,
		createdAt:     createdAt,
		closedAt:      closedAt,
	}
}
