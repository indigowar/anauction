// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package data

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteItem = `-- name: DeleteItem :one
DELETE FROM items
WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteItem(ctx context.Context, id pgtype.UUID) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, deleteItem, id)
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, deleteUser, id)
	err := row.Scan(&id)
	return id, err
}

const getByEmail = `-- name: GetByEmail :one
SELECT id, name, email, password, image FROM users
WHERE email = $1
`

func (q *Queries) GetByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Image,
	)
	return i, err
}

const getByID = `-- name: GetByID :one
SELECT id, name, email, password, image FROM users
WHERE id = $1
`

func (q *Queries) GetByID(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Image,
	)
	return i, err
}

const getItemByID = `-- name: GetItemByID :one
SELECT id, owner, name, image, description, start_price, created_at, closed_at FROM items
WHERE id = $1
`

func (q *Queries) GetItemByID(ctx context.Context, id pgtype.UUID) (Item, error) {
	row := q.db.QueryRow(ctx, getItemByID, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Image,
		&i.Description,
		&i.StartPrice,
		&i.CreatedAt,
		&i.ClosedAt,
	)
	return i, err
}

const getItemsByOwner = `-- name: GetItemsByOwner :many
SELECT id, owner, name, image, description, start_price, created_at, closed_at FROM items
WHERE owner = $1
`

func (q *Queries) GetItemsByOwner(ctx context.Context, owner pgtype.UUID) ([]Item, error) {
	rows, err := q.db.Query(ctx, getItemsByOwner, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Name,
			&i.Image,
			&i.Description,
			&i.StartPrice,
			&i.CreatedAt,
			&i.ClosedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertItem = `-- name: InsertItem :one
INSERT INTO items(
	id,
	owner,
	name,
	image,
	description,
	start_price,
	created_at,
	closed_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, owner, name, image, description, start_price, created_at, closed_at
`

type InsertItemParams struct {
	ID          pgtype.UUID
	Owner       pgtype.UUID
	Name        string
	Image       string
	Description string
	StartPrice  float64
	CreatedAt   pgtype.Timestamp
	ClosedAt    pgtype.Timestamp
}

func (q *Queries) InsertItem(ctx context.Context, arg InsertItemParams) (Item, error) {
	row := q.db.QueryRow(ctx, insertItem,
		arg.ID,
		arg.Owner,
		arg.Name,
		arg.Image,
		arg.Description,
		arg.StartPrice,
		arg.CreatedAt,
		arg.ClosedAt,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Name,
		&i.Image,
		&i.Description,
		&i.StartPrice,
		&i.CreatedAt,
		&i.ClosedAt,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO users(id, name, email, password)
VALUES ($1, $2, $3, $4)
`

type InsertUserParams struct {
	ID       pgtype.UUID
	Name     string
	Email    string
	Password string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.Exec(ctx, insertUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
	)
	return err
}

const updateItem = `-- name: UpdateItem :one
UPDATE items
SET
	owner = $2,
	name = $3,
	image = $4,
	description = $5,
	start_price = $6,
	created_at = $7,
	closed_at = $8
WHERE id = $1
RETURNING id
`

type UpdateItemParams struct {
	ID          pgtype.UUID
	Owner       pgtype.UUID
	Name        string
	Image       string
	Description string
	StartPrice  float64
	CreatedAt   pgtype.Timestamp
	ClosedAt    pgtype.Timestamp
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, updateItem,
		arg.ID,
		arg.Owner,
		arg.Name,
		arg.Image,
		arg.Description,
		arg.StartPrice,
		arg.CreatedAt,
		arg.ClosedAt,
	)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, password = $4, image = $5
WHERE id = $1
RETURNING id, name, email, password, image
`

type UpdateUserParams struct {
	ID       pgtype.UUID
	Name     string
	Email    string
	Password string
	Image    pgtype.Text
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Image,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Image,
	)
	return i, err
}
