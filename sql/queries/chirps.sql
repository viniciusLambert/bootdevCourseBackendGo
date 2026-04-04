-- name: CreateChirpy :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2
)
RETURNING *;

-- name: GetChirps :many
SELECT  * FROM chirps
order By created_at;

-- name: GetChirpsByUserID :many
SELECT * FROM chirps
where user_id = $1 
order by created_at;

-- name: GetChirpByID :one
SELECT  * FROM chirps
where id = $1;

-- name: DeleteChirpById :exec
delete from chirps
where id = $1;
