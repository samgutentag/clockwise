package main

import (
	"log/slog"

	"github.com/alecsavvy/clockwise/db"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// struct that implements the comet bft application interface
type App struct{
	logger *slog.Logger
	pool *pgxpool.Pool
	queries *db.Queries

	onGoingTx pgx.Tx
}

var _ abcitypes.Application = (*App)(nil)
