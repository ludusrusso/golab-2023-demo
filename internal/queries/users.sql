-- name: ListUsers :many
SELECT * FROM users 
	WHERE @match_labels::VARCHAR[] IS NULL OR labels @> @match_labels
	LIMIT sqlc.arg('limit') 
	OFFSET sqlc.arg('offset');

-- name: CountUsers :one
SELECT COUNT(*) FROM users WHERE @match_labels::VARCHAR[] IS NULL OR labels @> @match_labels;

-- name: CreateUser :one
INSERT INTO users (name, labels) VALUES (@name, @labels) RETURNING *;

-- name: DeleteUser :one
DELETE FROM users WHERE id = @id RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = @id LIMIT 1;