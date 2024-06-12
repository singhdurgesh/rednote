-- +goose Up
-- +goose StatementBegin
CREATE TABLE settings (
    id BIGSERIAL PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    title varchar(100),
    description text,
    key VARCHAR(100),
    value text,
    type int, -- Enum
    deleted_at timestamp with time zone
);
CREATE UNIQUE INDEX idx_settings_key ON settings(key);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.settings;
-- +goose StatementEnd
