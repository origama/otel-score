package example

import (
	"context"

	"otelscore/internal/plugins"
	"otelscore/internal/repository"
)

// RuleExample is a simple rule implementation that always passes.
type RuleExample struct{}

// New returns a new RuleExample.
func New() *RuleExample { return &RuleExample{} }

func (r *RuleExample) ID() string                  { return "EXAMPLE-001" }
func (r *RuleExample) Description() string         { return "Example rule that always passes" }
func (r *RuleExample) Impact() plugins.ImpactLevel { return plugins.ImpactLow }

// Check demonstrates using the repository manager to query telemetry.
func (r *RuleExample) Check(ctx context.Context, mgr *repository.Manager) (bool, error) {
	if repo := mgr.Default(); repo != nil {
		_, _ = repo.Query(ctx, "example_query")
	}
	return true, nil
}
