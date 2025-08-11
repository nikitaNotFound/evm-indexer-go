-- name: AddUniswapV3Pool :exec
INSERT INTO uniswap_v3_pools (address, token0, token1, fee, tick_spacing)
VALUES ($1, $2, $3, $4, $5);
