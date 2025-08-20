-- blocks
CREATE TABLE IF NOT EXISTS blocks (
    number BIGINT PRIMARY KEY,
    hash VARCHAR(255) NOT NULL,
    gas_price VARCHAR(255) NOT NULL,
    total_fees VARCHAR(255) NOT NULL,
    total_gas BIGINT NOT NULL,
    burnt_fees VARCHAR(255) NOT NULL,
    fee_recipient VARCHAR(255) NOT NULL,
    timestamp BIGINT NOT NULL
); 

-- raw txs
CREATE TABLE IF NOT EXISTS raw_txs (
    hash VARCHAR(255) PRIMARY KEY,
    from_address VARCHAR(255) NOT NULL,
    to_address VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    timestamp BIGINT NOT NULL,
    block_number BIGINT NOT NULL,
    input_data VARCHAR(255) NOT NULL,
    gas_used VARCHAR(255) NOT NULL,
    gas_price VARCHAR(255) NOT NULL,
    gas_limit VARCHAR(255) NOT NULL,
    max_priority_fee VARCHAR(255) NOT NULL,
    max_fee VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_raw_txs_from_address ON raw_txs (from_address);
CREATE INDEX IF NOT EXISTS idx_raw_txs_to_address ON raw_txs (to_address);
CREATE INDEX IF NOT EXISTS idx_raw_txs_block_number ON raw_txs (block_number);

-- uniswap v2 pools
CREATE TABLE IF NOT EXISTS uniswap_v2_pools (
    address VARCHAR(255) PRIMARY KEY,
    token0 VARCHAR(255) NOT NULL,
    token1 VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_uniswap_v2_pools_token0 ON uniswap_v2_pools (token0);
CREATE INDEX IF NOT EXISTS idx_uniswap_v2_pools_token1 ON uniswap_v2_pools (token1);

-- uniswap v3 pools
CREATE TABLE IF NOT EXISTS uniswap_v3_pools (
    address VARCHAR(255) PRIMARY KEY,
    token0 VARCHAR(255) NOT NULL,
    token1 VARCHAR(255) NOT NULL,
    fee VARCHAR(255) NOT NULL,
    tick_spacing VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_uniswap_v3_pools_token0 ON uniswap_v3_pools (token0);
CREATE INDEX IF NOT EXISTS idx_uniswap_v3_pools_token1 ON uniswap_v3_pools (token1);

-- pancake v2 pools
CREATE TABLE IF NOT EXISTS pancake_v2_pools (
    address VARCHAR(255) PRIMARY KEY,
    token0 VARCHAR(255) NOT NULL,
    token1 VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_pancake_v2_pools_token0 ON pancake_v2_pools (token0);
CREATE INDEX IF NOT EXISTS idx_pancake_v2_pools_token1 ON pancake_v2_pools (token1);

-- pancake v3 pools
CREATE TABLE IF NOT EXISTS pancake_v3_pools (
    address VARCHAR(255) PRIMARY KEY,
    token0 VARCHAR(255) NOT NULL,
    token1 VARCHAR(255) NOT NULL,
    fee VARCHAR(255) NOT NULL,
    tick_spacing VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_pancake_v3_pools_token0 ON pancake_v3_pools (token0);
CREATE INDEX IF NOT EXISTS idx_pancake_v3_pools_token1 ON pancake_v3_pools (token1);

-- sushiswap v2 pools
CREATE TABLE IF NOT EXISTS sushiswap_v2_pools (
    address VARCHAR(255) PRIMARY KEY,
    token0 VARCHAR(255) NOT NULL,
    token1 VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_sushiswap_v2_pools_token0 ON sushiswap_v2_pools (token0);
CREATE INDEX IF NOT EXISTS idx_sushiswap_v2_pools_token1 ON sushiswap_v2_pools (token1);

-- sushiswap v3 pools
CREATE TABLE IF NOT EXISTS sushiswap_v3_pools (
    address VARCHAR(255) PRIMARY KEY,
    token0 VARCHAR(255) NOT NULL,
    token1 VARCHAR(255) NOT NULL,
    fee VARCHAR(255) NOT NULL,
    tick_spacing VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_sushiswap_v3_pools_token0 ON sushiswap_v3_pools (token0);
CREATE INDEX IF NOT EXISTS idx_sushiswap_v3_pools_token1 ON sushiswap_v3_pools (token1);


