// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reads.sql

package db

import (
	"context"
)

const getTxResult = `-- name: GetTxResult :one
select rowid, block_id, index, created_at, tx_hash, tx_result from tx_results where tx_hash = $1
`

func (q *Queries) GetTxResult(ctx context.Context, txHash string) (TxResult, error) {
	row := q.db.QueryRow(ctx, getTxResult, txHash)
	var i TxResult
	err := row.Scan(
		&i.Rowid,
		&i.BlockID,
		&i.Index,
		&i.CreatedAt,
		&i.TxHash,
		&i.TxResult,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
select handle, tx_hash
from users
order by handle
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.Handle, &i.TxHash); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
