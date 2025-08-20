-- name: GetAllPancakeV3Pools :many
SELECT * FROM pancake_v3_pools
ORDER BY address
LIMIT $1 OFFSET $2;

-- name: CountPancakeV3Pools :one
SELECT COUNT(*) FROM pancake_v3_pools;

-- name: AddPancakeV3Pool :exec
INSERT INTO pancake_v3_pools (address, token0, token1, fee, tick_spacing)
VALUES ($1, $2, $3, $4, $5);