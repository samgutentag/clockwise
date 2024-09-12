package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/Khan/genqlient/generate"
	"github.com/alecsavvy/clockwise/db"
	"github.com/cometbft/cometbft/rpc/client/local"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
)

func run() error {
	// app level context
	ctx := context.Background()

	nodeID := os.Getenv("nodeId")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)).With("node", nodeID)

	// logger setup
	logger.Info("good morning!")

	// db setup
	pgConnectionString := os.Getenv("pgConnectionString")

	err := db.RunMigrations(pgConnectionString)
	if err != nil {
		return fmt.Errorf("could not complete database migrations: %v", err)
	}

	pool, err := pgxpool.New(ctx, pgConnectionString)
	if err != nil {
		return fmt.Errorf("failure to create db pool: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	node, err := NewNode("./cmt-home", pgConnectionString, &App{
		pool: pool,
		logger: logger,
		queries: queries,
	})
	if err != nil {
		return err
	}
	defer func() {
			node.Stop()
			node.Wait()
	}()

	rpc := local.New(node)

	api := &ApiServer {
		rpc: rpc,
		queries: queries,
		node: nodeID,
	}

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		logger.Info("starting node")

		node.Start()

		c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    <-c
		
		return nil
	})

	g.Go(func() error {
		return api.Start()
	})

	g.Go(func() error {
		time.Sleep(10 * time.Second)
		simulation(logger, nodeID, rpc)
		return nil
	})

	return g.Wait()
}

func main() {
	if err := run(); err != nil {
		fmt.Println("fatal error: ", err)
	}
	fmt.Println("run done")
}
