-- name: GetBlockByNumber :one
SELECT * FROM blocks WHERE number = $1;

-- name: AddBlock :exec
INSERT INTO blocks (number, hash, timestamp) VALUES ($1, $2, $3);