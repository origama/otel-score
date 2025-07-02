package plugins

import (
	"context"

	"otelscore/internal/repository"
)

// Registry holds registered rule checkers.
type Registry struct {
	checkers []RuleChecker
}

// NewRegistry creates an empty Registry.
func NewRegistry() *Registry {
	return &Registry{}
}

// Register adds a new RuleChecker to the registry.
func (r *Registry) Register(rc RuleChecker) {
	r.checkers = append(r.checkers, rc)
}

// RunAll executes all registered checkers with the given Repository Manager.
func (r *Registry) RunAll(ctx context.Context, mgr *repository.Manager) []Result {
	var results []Result
	for _, c := range r.checkers {
		passed, err := c.Check(ctx, mgr)
		results = append(results, Result{
			ID:          c.ID(),
			Description: c.Description(),
			Impact:      c.Impact(),
			Passed:      passed,
			Error:       err,
		})
	}
	return results
}
