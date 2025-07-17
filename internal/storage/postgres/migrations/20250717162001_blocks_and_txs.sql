CREATE TABLE blocks (
    number BIGINT PRIMARY KEY,
    hash VARCHAR(255) NOT NULL,
    timestamp BIGINT NOT NULL
);