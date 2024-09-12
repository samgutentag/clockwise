package main

import (
	"context"

	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	"github.com/jackc/pgx/v5"
)

// CheckTx implements types.Application.
func (a *App) CheckTx(ctx context.Context, req *v1.CheckTxRequest) (*v1.CheckTxResponse, error) {
	a.logger.Info("checking tx")
	return &v1.CheckTxResponse{Code: v1.CodeTypeOK}, nil
}

// PrepareProposal implements types.Application.
func (a *App) PrepareProposal(ctx context.Context, req *v1.PrepareProposalRequest) (*v1.PrepareProposalResponse, error) {
	a.logger.Info("i am preparing a new block")
	return &v1.PrepareProposalResponse{Txs: req.GetTxs()}, nil
}

// ProcessProposal implements types.Application.
func (a *App) ProcessProposal(ctx context.Context, req *v1.ProcessProposalRequest) (*v1.ProcessProposalResponse, error) {
	a.logger.Info("i am processing a new block proposal")
	return &v1.ProcessProposalResponse{Status: v1.PROCESS_PROPOSAL_STATUS_ACCEPT}, nil
}

// FinalizeBlock implements types.Application.
func (a *App) FinalizeBlock(ctx context.Context, req *v1.FinalizeBlockRequest) (*v1.FinalizeBlockResponse, error) {
	a.logger.Info("i am finalizing a new block")

	finalizeDbTx, err := a.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		a.logger.Error("could not open db tx", "error", err)
		return &v1.FinalizeBlockResponse{}, nil
	}

	// set ongoing tx for use in commit
	a.onGoingTx = finalizeDbTx


	return &v1.FinalizeBlockResponse{}, nil
}

// Commit implements types.Application.
func (a *App) Commit(ctx context.Context, req *v1.CommitRequest) (*v1.CommitResponse, error) {

	if a.onGoingTx == nil {
		return &v1.CommitResponse{}, nil
	}

	err := a.onGoingTx.Commit(ctx)
	if err != nil {
		a.logger.Error("could not ")
	}
	panic("unimplemented")
}
