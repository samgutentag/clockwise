package main

import (
	"context"

	"github.com/alecsavvy/clockwise/proto_gen"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/proto"
)

// CheckTx implements types.Application.
func (a *App) CheckTx(ctx context.Context, req *v1.CheckTxRequest) (*v1.CheckTxResponse, error) {
	a.logger.Info("checking tx")

	// TODO: check if user exists first

	return &v1.CheckTxResponse{Code: v1.CodeTypeOK}, nil
}

// PrepareProposal implements types.Application.
func (a *App) PrepareProposal(ctx context.Context, req *v1.PrepareProposalRequest) (*v1.PrepareProposalResponse, error) {
	a.logger.Info("i am preparing a new block", "height", req.Height)

	// TODO: reorder txs so they make sense

	return &v1.PrepareProposalResponse{Txs: req.GetTxs()}, nil
}

// ProcessProposal implements types.Application.
func (a *App) ProcessProposal(ctx context.Context, req *v1.ProcessProposalRequest) (*v1.ProcessProposalResponse, error) {
	a.logger.Info("i am processing a new block proposal", "height", req.Height)

	// TODO: validate txs

	return &v1.ProcessProposalResponse{Status: v1.PROCESS_PROPOSAL_STATUS_ACCEPT}, nil
}

// FinalizeBlock implements types.Application.
func (a *App) FinalizeBlock(ctx context.Context, req *v1.FinalizeBlockRequest) (*v1.FinalizeBlockResponse, error) {
	a.logger.Info("i am finalizing a new block", "height", req.Height, "txs", len(req.Txs))

	finalizeDbTx, err := a.pool.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		a.logger.Error("could not open db tx", "error", err)
		return &v1.FinalizeBlockResponse{}, nil
	}

	// set ongoing tx for use in commit
	a.onGoingTx = finalizeDbTx

	var txs = make([]*abcitypes.ExecTxResult, len(req.Txs))

	for i, tx := range req.Txs {
		err := func() error {
			manageEntity := &proto_gen.ManageEntity{}
			err := proto.Unmarshal(tx, manageEntity)
			if err != nil {
				return err
			}

			switch manageEntity.Message.(type) {
			case *proto_gen.ManageEntity_UserCreate:
			case *proto_gen.ManageEntity_TrackCreate:
			}

			return nil
		}()

		if err != nil {
			a.logger.Error("error with finalizing tx", "error", err)
			txs[i] = &abcitypes.ExecTxResult{Code: 1}
		} else {
			txs[i] = &abcitypes.ExecTxResult{Code: v1.CodeTypeOK}
		}
	}

	return &v1.FinalizeBlockResponse{}, nil
}

// Commit implements types.Application.
func (a *App) Commit(ctx context.Context, req *v1.CommitRequest) (*v1.CommitResponse, error) {

	if a.onGoingTx == nil {
		return &v1.CommitResponse{}, nil
	}

	err := a.onGoingTx.Commit(ctx)
	if err != nil {
		a.logger.Error("could not commit txs")
	}
	return &v1.CommitResponse{}, nil
}
