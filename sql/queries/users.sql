-- name: CreateUser :one
INSERT INTO users (id, email, hashed_password, created_at, updated_at)
VALUES (
  gen_random_uuid(),
  $1,
  "unset",
  NOW(),
  NOW() 
)
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE id=id;
