package main

import (
	"context"
	"log/slog"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
)

// struct that implements the comet bft application interface
type App struct{
	logger *slog.Logger
}

// ApplySnapshotChunk implements types.Application.
func (a *App) ApplySnapshotChunk(ctx context.Context, req *v1.ApplySnapshotChunkRequest) (*v1.ApplySnapshotChunkResponse, error) {
	panic("unimplemented")
}

// CheckTx implements types.Application.
func (a *App) CheckTx(ctx context.Context, req *v1.CheckTxRequest) (*v1.CheckTxResponse, error) {
	panic("unimplemented")
}

// Commit implements types.Application.
func (a *App) Commit(ctx context.Context, req *v1.CommitRequest) (*v1.CommitResponse, error) {
	panic("unimplemented")
}

// ExtendVote implements types.Application.
func (a *App) ExtendVote(ctx context.Context, req *v1.ExtendVoteRequest) (*v1.ExtendVoteResponse, error) {
	panic("unimplemented")
}

// FinalizeBlock implements types.Application.
func (a *App) FinalizeBlock(ctx context.Context, req *v1.FinalizeBlockRequest) (*v1.FinalizeBlockResponse, error) {
	panic("unimplemented")
}

// Info implements types.Application.
func (a *App) Info(ctx context.Context, req *v1.InfoRequest) (*v1.InfoResponse, error) {
	a.logger.Info("info call")
	return &v1.InfoResponse{}, nil
}

// InitChain implements types.Application.
func (a *App) InitChain(ctx context.Context, req *v1.InitChainRequest) (*v1.InitChainResponse, error) {
	return &v1.InitChainResponse{}, nil
}

// ListSnapshots implements types.Application.
func (a *App) ListSnapshots(ctx context.Context, req *v1.ListSnapshotsRequest) (*v1.ListSnapshotsResponse, error) {
	panic("unimplemented")
}

// LoadSnapshotChunk implements types.Application.
func (a *App) LoadSnapshotChunk(ctx context.Context, req *v1.LoadSnapshotChunkRequest) (*v1.LoadSnapshotChunkResponse, error) {
	panic("unimplemented")
}

// OfferSnapshot implements types.Application.
func (a *App) OfferSnapshot(ctx context.Context, req *v1.OfferSnapshotRequest) (*v1.OfferSnapshotResponse, error) {
	panic("unimplemented")
}

// PrepareProposal implements types.Application.
func (a *App) PrepareProposal(ctx context.Context, req *v1.PrepareProposalRequest) (*v1.PrepareProposalResponse, error) {
	panic("unimplemented")
}

// ProcessProposal implements types.Application.
func (a *App) ProcessProposal(ctx context.Context, req *v1.ProcessProposalRequest) (*v1.ProcessProposalResponse, error) {
	panic("unimplemented")
}

// Query implements types.Application.
func (a *App) Query(ctx context.Context, req *v1.QueryRequest) (*v1.QueryResponse, error) {
	panic("unimplemented")
}

// VerifyVoteExtension implements types.Application.
func (a *App) VerifyVoteExtension(ctx context.Context, req *v1.VerifyVoteExtensionRequest) (*v1.VerifyVoteExtensionResponse, error) {
	panic("unimplemented")
}

var _ abcitypes.Application = (*App)(nil)
