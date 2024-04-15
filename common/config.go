package common

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

type NodeType string

const (
	Observer NodeType = "observer"
	Sealer   NodeType = "sealer"
	Archiver NodeType = "archiver"
)

type Config struct {
	/** Network Config */
	NetworkID                 string
	NetworkTransactionCeiling uint32
	NetworkBlockMineMaximumMS uint32

	/** Node Config */
	NodeEndpoint       string
	NodeBootstrapPeers []string
	NodeType           NodeType

	/** Database Config */
	DatabaseBlockPersistenceThreshold uint32
	DatabasePruningIntervalMS         uint32
	DatabaseFilePath                  string

	/** Peer Config */
	PeerMaximumConnected      uint32
	PeerReconnectIntervalMS   uint32
	PeerHealthCheckIntervalMS uint32

	/** UI Config */
	UIEndpoint string
}

func DefaultConfig() *Config {
	defaultEndpoint := "localhost:8080"
	return &Config{
		NetworkID:                         "dev",
		NetworkTransactionCeiling:         1000,
		NetworkBlockMineMaximumMS:         5000,
		NodeEndpoint:                      defaultEndpoint,
		NodeBootstrapPeers:                []string{"localhost:8081", "localhost:8082"},
		NodeType:                          "sealer",
		DatabaseBlockPersistenceThreshold: 15000,
		DatabasePruningIntervalMS:         60000,
		DatabaseFilePath:                  ":memory:",
		PeerMaximumConnected:              0,
		PeerReconnectIntervalMS:           60000,
		PeerHealthCheckIntervalMS:         5000,
		UIEndpoint:                        fmt.Sprintf("%s0", defaultEndpoint),
	}
}

func ReadConfig(path string) (*Config, error) {
	config := DefaultConfig()
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	if !IsValidNodeType(config.NodeType) {
		return nil, errors.New(fmt.Sprintf("invalid node type %s", config.NodeType))
	}
	return config, nil
}

func IsValidNodeType(nodeType NodeType) bool {
	switch NodeType(nodeType) {
	case Observer, Sealer, Archiver:
		return true
	}
	return false
}
