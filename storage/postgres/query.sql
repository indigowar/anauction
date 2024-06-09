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
