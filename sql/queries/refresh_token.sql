-- name: CreateRefreshToken :one
insert into refresh_tokens(token, created_at, updated_at, user_id, expired_at, revoked_at)
values (
  $1,
  NOW(),
  NOW(),
  $2,
   NOW() +  interval  '60 days',
  null
)
returning *;


-- name: GetRefreshTokenByToken :one
select * from refresh_tokens
where token = $1;

-- name: RevokeToken :exec
update refresh_tokens
set revoked_at = NOW(),
updated_at = NOW()
where token = $1;
