-- name: GetAllSushiSwapV2Pools :many
SELECT * FROM sushiswap_v2_pools
ORDER BY address
LIMIT $1 OFFSET $2;

-- name: CountSushiSwapV2Pools :one
SELECT COUNT(*) FROM sushiswap_v2_pools;

-- name: AddSushiSwapV2Pool :exec
INSERT INTO sushiswap_v2_pools (address, token0, token1)
VALUES ($1, $2, $3);