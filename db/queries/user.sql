-- name: CreateUser :one
INSERT INTO users (id, first_name, last_name, nickname, password, email, country)
VALUES (gen_random_uuid(), $1, $2, $3, crypt($4, gen_salt('bf')), $5, $6)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET first_name = COALESCE($2, first_name),
    last_name = COALESCE($3, last_name),
    nickname = COALESCE($4, nickname),
    password = COALESCE(crypt($5, gen_salt('bf')), password),
    email = COALESCE($6, email),
    country = COALESCE($7, country),
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
WHERE ($1::TEXT IS NULL OR country = $1)  -- Filter by country if provided
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;
