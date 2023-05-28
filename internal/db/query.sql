-- name: InsertTx :exec
INSERT INTO tx (id, date_posted, memo) VALUES (null, ?, ?)
RETURNING *;

-- name: GetTxs :many
SELECT * FROM tx ORDER BY id;
