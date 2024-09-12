package main

import (
	"context"
	"fmt"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/proxy"
	"github.com/spf13/viper"
)

func NewNode(homeDir string, psqlConn string, app abcitypes.Application) (*node.Node, error) {
	config := cfg.DefaultConfig()
	config.SetRoot(homeDir)
	viper.SetConfigFile(fmt.Sprintf("%s/%s", homeDir, "config/config.toml"))

	// replace this with go config?
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	if err := config.ValidateBasic(); err != nil {
		return nil, err
	}

	// load priv key, this will be replaced by SP derived key
	pv := privval.LoadFilePV(
		config.PrivValidatorKeyFile(),
		config.PrivValidatorStateFile(),
	)

	nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, err
	}

	// set up postgres indexing instead of kv store
	config.TxIndex.Indexer = "psql"
	config.TxIndex.PsqlConn = psqlConn

	node, err := node.NewNode(
		context.Background(),
		config,
		pv,
		nodeKey,
		proxy.NewLocalClientCreator(app),
		node.DefaultGenesisDocProviderFunc(config),
		cfg.DefaultDBProvider,
		node.DefaultMetricsProvider(config.Instrumentation),
		NoopLogger{},
	)

	if err != nil {
		return nil, err
	}

	return node, nil
}
