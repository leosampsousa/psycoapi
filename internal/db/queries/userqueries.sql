-- name: SaveUser :exec
INSERT INTO users (first_name, last_name, username, hashed_password)
VALUES ($1, $2, $3, $4);

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;