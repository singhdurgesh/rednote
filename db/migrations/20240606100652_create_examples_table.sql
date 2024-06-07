-- +goose Up
-- +goose StatementBegin
CREATE TABLE examples (
    id BIGSERIAL PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name VARCHAR,
    status VARCHAR,
    deleted_at timestamp with time zone
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.examples;
-- +goose StatementEnd
