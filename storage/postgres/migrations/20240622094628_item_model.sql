-- +goose Up
-- +goose StatementBegin
CREATE TABLE items(
	id UUID PRIMARY KEY,
	
	owner UUID NOT NULL,
	FOREIGN KEY(owner) REFERENCES users(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,
	
	name VARCHAR(255) NOT NULL,
	image VARCHAR(1024) NOT NULL,
	description VARCHAR(1024) NOT NULL,

	start_price FLOAT NOT NULL
		CHECK (start_price >= 0),
	
	created_at TIMESTAMP,
	closed_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
