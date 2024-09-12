package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/Khan/genqlient/generate"
	"github.com/alecsavvy/clockwise/db"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
)

func run() error {
	// app level context
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)).With("node", os.Getenv("nodeId"))

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

	node, err := NewNode("./cmt-home", pgConnectionString, &App{
		logger: logger,
	})
	if err != nil {
		return err
	}
	defer func() {
			node.Stop()
			node.Wait()
	}()

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		logger.Info("starting node")

		node.Start()

		c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    <-c
		
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
