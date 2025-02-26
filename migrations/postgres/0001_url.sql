-- +goose Up
CREATE TABLE IF NOT EXISTS url (
    id SERIAL PRIMARY KEY,
    original_url VARCHAR(255) UNIQUE NOT NULL,
    code VARCHAR(8) NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS url;