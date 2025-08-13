-- name: GetBlockByNumber :one
SELECT * FROM blocks WHERE number = $1;

-- name: AddBlock :exec
INSERT INTO blocks (number, hash, timestamp, gas_price, total_fees, total_gas, burnt_fees, fee_recipient)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);