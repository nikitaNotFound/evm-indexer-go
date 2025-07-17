-- name: GetBlockByNumber :one
SELECT * FROM blocks WHERE number = $1;