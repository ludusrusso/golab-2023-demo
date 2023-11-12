-- name: ListUsers :many
SELECT * FROM users LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: CreateUser :one
INSERT INTO users (name) VALUES (@name) RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = @id RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = @id LIMIT 1;