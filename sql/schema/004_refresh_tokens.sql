-- +goose Up
create table refresh_tokens(
  token text primary key,
  created_at timestamp not  null,
  updated_at timestamp not null,
  user_id uuid not null  references users(id) on delete cascade,
  expired_at timestamp,
  revoked_at timestamp
);

-- +goose Down
drop table refresh_tokens;
