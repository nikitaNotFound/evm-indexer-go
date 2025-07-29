CREATE TABLE IF NOT EXISTS blocks (
    number BIGINT PRIMARY KEY,
    hash VARCHAR(255) NOT NULL,
    timestamp BIGINT NOT NULL
); 

CREATE TABLE IF NOT EXISTS raw_txs (
    hash VARCHAR(255) PRIMARY KEY,
    from_address VARCHAR(255) NOT NULL,
    to_address VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    timestamp BIGINT NOT NULL,
    block_number BIGINT NOT NULL,
    input_data VARCHAR(255) NOT NULL,
    gas_used VARCHAR(255) NOT NULL,
    max_gas_price VARCHAR(255) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_raw_txs_from_address ON raw_txs (from_address);
CREATE INDEX IF NOT EXISTS idx_raw_txs_to_address ON raw_txs (to_address);
CREATE INDEX IF NOT EXISTS idx_raw_txs_block_number ON raw_txs (block_number);