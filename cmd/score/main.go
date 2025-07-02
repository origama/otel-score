package main

import (
	"context"
	"fmt"

	"otelscore/internal/config"
	"otelscore/internal/logging"
	"otelscore/internal/metrics"
	"otelscore/internal/plugins"
	"otelscore/internal/repository"
	"otelscore/plugins/example"
)

func main() {
	ctx := context.Background()

	cfg := config.Load()
	logger := logging.New()

	metrics.Init()

	repoMgr := repository.NewManager(cfg.Repositories)

	registry := plugins.NewRegistry()
	registry.Register(example.New())

	results := registry.RunAll(ctx, repoMgr)

	for _, r := range results {
		if r.Error != nil {
			logger.Println("error running", r.ID, r.Error)
			continue
		}
		logger.Printf("Rule %s passed: %v", r.ID, r.Passed)
	}

	if repo := repoMgr.Default(); repo != nil {
		fmt.Println("Default repo configured")
	}
}
