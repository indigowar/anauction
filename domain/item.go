package domain

import "github.com/google/uuid"

type Item struct {
	ID    uuid.UUID
	Owner uuid.UUID
	Name  string
}
