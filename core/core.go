/* the client that external modules use like grpc and clis */
package core

import (
	"fmt"

	"github.com/alecsavvy/clockwise/core/db"
	"github.com/alecsavvy/clockwise/utils"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	cfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/node"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	"github.com/cometbft/cometbft/proxy"
	"github.com/cometbft/cometbft/rpc/client/local"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type Core struct {
	logger    *utils.Logger
	rpc       *local.Local
	db        *db.Queries
	pool      *pgxpool.Pool
	pubsub    *Pubsub
	currentTx pgx.Tx
}

var _ abcitypes.Application = (*Core)(nil)

func NewCore(logger *utils.Logger, pool *pgxpool.Pool) *Core {
	return &Core{
		logger:    logger,
		db:        db.New(pool),
		pool:      pool,
		pubsub:    NewPubsub(),
		currentTx: nil,
	}
}

func NewNode(logger *utils.Logger, homeDir string, app abcitypes.Application) (*node.Node, error) {
	config := cfg.DefaultConfig()
	config.SetRoot(homeDir)
	viper.SetConfigFile(fmt.Sprintf("%s/%s", homeDir, "config/config.toml"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, utils.AppError("Reading config", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, utils.AppError("Decoding config", err)
	}
	if err := config.ValidateBasic(); err != nil {
		return nil, utils.AppError("Chain config validation", err)
	}

	pv := privval.LoadFilePV(
		config.PrivValidatorKeyFile(),
		config.PrivValidatorStateFile(),
	)

	nodeKey, err := p2p.LoadNodeKey(config.NodeKeyFile())
	if err != nil {
		return nil, utils.AppError("Error loading p2p key", err)
	}

	node, err := node.NewNode(
		config,
		pv,
		nodeKey,
		proxy.NewLocalClientCreator(app),
		node.DefaultGenesisDocProviderFunc(config),
		cfg.DefaultDBProvider,
		node.DefaultMetricsProvider(config.Instrumentation),
		logger,
	)

	if err != nil {
		return nil, utils.AppError("Error initializing node", err)
	}

	return node, nil
}

func (c *Core) Run(node *node.Node) error {
	c.rpc = local.New(node)

	// start all the core processes
	// pubsub
	// eth indexer maybe? idk

	return nil
}
