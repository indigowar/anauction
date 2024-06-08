-- name: InsertUser :exec
INSERT INTO users(id, name, email, password)
VALUES ($1, $2, $3, $4);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetByID :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET name = $2, email = $3, password = $4, image = $5
WHERE id = $1;
