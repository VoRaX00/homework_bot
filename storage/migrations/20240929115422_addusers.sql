-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
   id uuid PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   study_group NUMERIC
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
