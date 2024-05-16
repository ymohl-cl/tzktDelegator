-- queries for sqlc

-- name: InsertDelegator :one
INSERT INTO tzkt.Delegation (
	delegation_date,
	delegator_address,
	block_hash,
	amount,
	block_state,
	external_id
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
) RETURNING *;

-- name: SearchDelegator :many
-- SearchDelegator query function
-- list parameters
-- @delegation_year::text year of the delegation
-- @limit_item::integer maximum number of items
SELECT * FROM tzkt.Delegation
WHERE
	(@delegation_year::integer = 0 OR @delegation_year = EXTRACT(YEAR FROM delegation_date))
ORDER BY delegation_date DESC
LIMIT
CASE
	WHEN
		(@limit_item::integer > 0)
	THEN @limit_item
END;

-- name: DeleteDelegator :many
-- DeleteDelegator query function
-- list parameters
-- @from_state::integer define the last valide block state
DELETE FROM tzkt.Delegation
WHERE
	block_state > @from_state::integer
RETURNING internal_id;