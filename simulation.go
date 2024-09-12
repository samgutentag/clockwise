package main

import "log/slog"

// function to simulate user simulation
func simulation(logger *slog.Logger, node string) error {
	if node != "1" {
		logger.Info("not simulation node, goodbye")
		return nil
	}

	logger.Info("starting user simulation")

	return nil
}
