-- name: CreateProduct :one
INSERT INTO products (name, category, location, quantity, description, price)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1
LIMIT 1;

-- name: UpdateProductById :one
UPDATE products
SET name = $1, category = $2, location = $3, quantity = $4, description = $5, price = $6, updated_at = NOW()
WHERE id = $7
RETURNING *;

-- name: DeleteProductById :exec
DELETE FROM products
WHERE id = $1;

-- name: GetAllProducts :many
SELECT * FROM products;