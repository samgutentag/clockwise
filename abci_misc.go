package main

import (
	"context"

	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
)

// VerifyVoteExtension implements types.Application.
func (a *App) VerifyVoteExtension(ctx context.Context, req *v1.VerifyVoteExtensionRequest) (*v1.VerifyVoteExtensionResponse, error) {
	panic("unimplemented")
}

// ExtendVote implements types.Application.
func (a *App) ExtendVote(ctx context.Context, req *v1.ExtendVoteRequest) (*v1.ExtendVoteResponse, error) {
	panic("unimplemented")
}

// Query implements types.Application.
func (a *App) Query(ctx context.Context, req *v1.QueryRequest) (*v1.QueryResponse, error) {
	panic("unimplemented")
}
