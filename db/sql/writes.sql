-- name: CreateUser :exec
insert into users (id, handle, bio, tx_hash)
values ($1, $2, $3, $4);
