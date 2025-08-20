-- name: GetAllUniswapV2Pools :many
SELECT * FROM uniswap_v2_pools
ORDER BY address
LIMIT $1 OFFSET $2;

-- name: CountUniswapV2Pools :one
SELECT COUNT(*) FROM uniswap_v2_pools;

-- name: AddUniswapV2Pool :exec
INSERT INTO uniswap_v2_pools (address, token0, token1) VALUES ($1, $2, $3);
