-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
	id UUID PRIMARY KEY,
	name VARCHAR(64) NOT NULL,
	email VARCHAR(128) UNIQUE NOT NULL,
	password VARCHAR(1024) NOT NULL,
	image TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
