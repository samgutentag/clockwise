package main

import (
	"context"
	"log/slog"

	"github.com/alecsavvy/clockwise/proto_gen"
	"github.com/cometbft/cometbft/rpc/client/local"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/protobuf/proto"
)

// function to simulate user simulation
func simulation(logger *slog.Logger, node string, rpc *local.Local) {
	ctx := context.Background()

	if node != "1" {
		logger.Info("not simulation node, goodbye")
		return
	}

	logger.Info("starting user simulation")

	createUser := proto_gen.ManageEntity{
		Signature: "i am a registered node, trust me plz",
		Message: &proto_gen.ManageEntity_UserCreate{
			UserCreate: &proto_gen.UserCreate{
				Handle: "lemonadejetpack",
			},
		},
	}

	createUserBytes, err := proto.Marshal(&createUser)
	if err != nil {
		logger.Error("error marshaling msg", "error", err)
		return
	}

	res, err := rpc.BroadcastTxSync(ctx, createUserBytes)
	if err != nil {
		logger.Error("error broadcasting tx", "error", err)
		return
	}

	spew.Dump(res)

	// TODO: implement POST api
	req := map[string]string{
		"name":  "Alec",
		"email": "alec@example.com",
	}

	postRes, err := SendPost("", req)
	if err != nil {
		logger.Error("problem with post", "error", err)
	}

	spew.Dump(postRes)
}
