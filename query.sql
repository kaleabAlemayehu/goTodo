-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;


-- name: GetUsers :many
SELECT * FROM users;


-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set username = $2,
  email = $3,
  password = $4
WHERE id = $1;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;