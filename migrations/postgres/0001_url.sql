-- +goose Up
CREATE TABLE url (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    code TEXT NOT NULL,
);


-- +goose Down
DROP TABLE url;