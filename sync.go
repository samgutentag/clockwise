package main

import (
	"context"

	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
)

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

// ApplySnapshotChunk implements types.Application.
func (a *App) ApplySnapshotChunk(ctx context.Context, req *v1.ApplySnapshotChunkRequest) (*v1.ApplySnapshotChunkResponse, error) {
	panic("unimplemented")
}
