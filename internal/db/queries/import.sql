-- name: CreateImport :one
INSERT INTO import (id, date_created, filename, balance_amount_num, balance_amount_den, date_as_of) VALUES (null, ?, ?, ?, ?, ?)
RETURNING *;
