-- name: GetUsers :many
select *
from users
order by handle;

-- name: GetTxResult :one
select * from tx_results where tx_hash = $1;
