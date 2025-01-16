-- name: SaveUser :exec
INSERT INTO users (first_name, last_name, username, hashed_password)
VALUES ($1, $2, $3, $4);

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByUsernameAndPassword :one
SELECT * FROM users
WHERE username = $1 AND hashed_password = $2 LIMIT 1;