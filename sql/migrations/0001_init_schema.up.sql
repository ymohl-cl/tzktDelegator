-- Schema definitions
CREATE SCHEMA IF NOT EXISTS tzkt;

-- Tables definitions
CREATE TABLE IF NOT EXISTS tzkt.Delegation (
	internal_id BIGSERIAL PRIMARY KEY,
	external_id BIGSERIAL NOT NULL,
	delegation_date TIMESTAMP NOT NULL,
	delegator_address TEXT NOT NULL,
	block_hash TEXT NOT NULL,
	amount TEXT NOT NULL,
	block_state BIGINT NOT NULL,
	UNIQUE (internal_id),
	UNIQUE (external_id)
);