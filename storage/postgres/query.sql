-- name: InsertUser :exec
INSERT INTO users(id, name, email, password)
VALUES ($1, $2, $3, $4);

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;

-- name: GetByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetByID :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, password = $4, image = $5
WHERE id = $1
RETURNING *;

-- name: InsertItem :one
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
RETURNING *;

-- name: GetItemByID :one
SELECT * FROM items
WHERE id = $1;

-- name: GetItemsByOwner :many
SELECT * FROM items
WHERE owner = $1;

-- name: DeleteItem :one
DELETE FROM items
WHERE id = $1
RETURNING id;

-- name: UpdateItem :one
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
RETURNING id;
