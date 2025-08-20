-- name: GetAllUniswapV3Pools :many
SELECT * FROM uniswap_v3_pools
ORDER BY address
LIMIT $1 OFFSET $2;

-- name: CountUniswapV3Pools :one
SELECT COUNT(*) FROM uniswap_v3_pools;

-- name: AddUniswapV3Pool :exec
INSERT INTO uniswap_v3_pools (address, token0, token1, fee, tick_spacing)
VALUES ($1, $2, $3, $4, $5);
