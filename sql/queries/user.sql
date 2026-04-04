-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2
)
RETURNING *;

-- name: UpdateUser :one
update users
set email = $2,
hashed_password = $3,
updated_at = NOW()
where id = $1
RETURNING *;


-- name: ResetUsersTable :exec
delete from Users;

-- name: GetUserByEmail :one
select * from users
where email = $1;
