-- name: GetAllPancakeV2Pools :many
SELECT * FROM pancake_v2_pools
ORDER BY address
LIMIT $1 OFFSET $2;

-- name: CountPancakeV2Pools :one
SELECT COUNT(*) FROM pancake_v2_pools;

-- name: AddPancakeV2Pool :exec
INSERT INTO pancake_v2_pools (address, token0, token1) VALUES ($1, $2, $3);