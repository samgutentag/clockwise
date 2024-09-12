-- name: CreateUser :exec
insert into users (handle, tx_hash)
values ($1, $2);
