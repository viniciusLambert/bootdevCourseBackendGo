-- +goose Up
CREATE TABLE users(
  id uuid PRIMARY KEY,
  created_at timestamp not  null,
  updated_at timestamp not null,
  email text not null,
  unique(email)
);

-- +goose Down
DROP TABLE users;
