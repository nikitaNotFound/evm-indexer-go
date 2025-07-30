-- name: AddUniswapV2Pool :exec
INSERT INTO uniswap_v2_pools (address, token0, token1) VALUES ($1, $2, $3);
