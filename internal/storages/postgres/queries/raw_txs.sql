-- name: AddRawTx :exec
INSERT INTO raw_txs (
    hash,
    from_address,
    to_address,
    value,
    timestamp,
    block_number,
    input_data,
    gas_used,
    max_gas_price
)
VALUES (
    sqlc.arg(hash)::VARCHAR(255),
    sqlc.arg(from_address)::VARCHAR(255),
    sqlc.arg(to_address)::VARCHAR(255),
    sqlc.arg(value)::VARCHAR(255),
    sqlc.arg(timestamp)::BIGINT,
    sqlc.arg(block_number)::BIGINT,
    sqlc.arg(input_data)::VARCHAR(255),
    sqlc.arg(gas_used)::VARCHAR(255),
    sqlc.arg(max_gas_price)::VARCHAR(255)
);
