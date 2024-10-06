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


-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;


-- name: GetTodos :many
SELECT * FROM todos;


-- name: ListTodos :many
SELECT * FROM todos
ORDER BY title;

-- name: CreateTodo :one
INSERT INTO todos (
user_id, title, content, starting_time, ending_time
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateTodo :exec
UPDATE todos
  set title = $2,
  content = $3,
  starting_time = $4,
  ending_time = $5
WHERE id = $1;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id = $1
RETURNING *;