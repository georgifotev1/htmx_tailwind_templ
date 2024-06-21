-- +goose Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(100) NOT NULL,
    location VARCHAR(100) NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now()),
    updated_at TIMESTAMP NOT NULL DEFAULT (now()),
    description TEXT,
    price NUMERIC(10, 2),
    is_active BOOLEAN DEFAULT TRUE
);

-- +goose Down
DROP TABLE products;