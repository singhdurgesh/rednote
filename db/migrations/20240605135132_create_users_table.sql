-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text,
    username text,
    password text,
    last_login_at timestamp with time zone,
    phone text,
    email text,
    email_verified bool,
    deleted_at timestamp with time zone
);

-- Indices -------------------------------------------------------

CREATE INDEX idx_users_deleted_at ON users(deleted_at timestamptz_ops);
CREATE UNIQUE INDEX idx_users_username ON users(username text_ops);
CREATE UNIQUE INDEX idx_users_phone ON users(phone text_ops);
CREATE UNIQUE INDEX idx_users_email ON users(email text_ops);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.users;
-- +goose StatementEnd
