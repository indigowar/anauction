// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package data

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       pgtype.UUID
	Name     string
	Email    string
	Password string
	Image    pgtype.Text
}
