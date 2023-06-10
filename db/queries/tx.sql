-- name: CreateTx :one
INSERT INTO tx (id, date_created, date_posted, memo, amount_num, amount_den, import_id) VALUES (NULL, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: ListTxs :many
SELECT * FROM tx ORDER BY id;

-- name: GetDuplicateTx :many
SELECT * FROM tx WHERE date_posted = ? AND memo = ? AND amount_num = ? AND amount_den = ?;
