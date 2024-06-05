package domain

import "github.com/google/uuid"

type Bid struct {
	ID    uuid.UUID
	Item  uuid.UUID
	Owner uuid.UUID
	Price float64
}
