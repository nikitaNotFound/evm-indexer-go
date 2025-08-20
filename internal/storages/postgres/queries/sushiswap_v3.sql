-- name: GetAllSushiSwapV3Pools :many
SELECT * FROM sushiswap_v3_pools
ORDER BY address
LIMIT $1 OFFSET $2;

-- name: CountSushiSwapV3Pools :one
SELECT COUNT(*) FROM sushiswap_v3_pools;

-- name: AddSushiSwapV3Pool :exec
INSERT INTO sushiswap_v3_pools (address, token0, token1, fee, tick_spacing)
VALUES ($1, $2, $3, $4, $5);