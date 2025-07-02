package repository

import (
	"fmt"

	"context"

	"otelscore/internal/config"
)

// Manager provides access to configured repositories.
type Manager struct {
	repos       map[string]Repository
	defaultRepo Repository
}

// NewManager initializes repositories defined in the configuration.
// Unknown repository types are ignored.
func NewManager(cfg []config.RepositoryConfig) *Manager {
	m := &Manager{repos: make(map[string]Repository)}
	for i, rc := range cfg {
		repo := newRepo(rc)
		if repo == nil {
			continue
		}
		name := rc.Name
		if name == "" {
			name = fmt.Sprintf("repo%d", i)
		}
		if i == 0 {
			m.defaultRepo = repo
		}
		m.repos[name] = repo
	}
	return m
}

func (m *Manager) Default() Repository {
	return m.defaultRepo
}

func (m *Manager) Get(name string) (Repository, bool) {
	r, ok := m.repos[name]
	return r, ok
}

func newRepo(cfg config.RepositoryConfig) Repository {
	switch cfg.Type {
	case config.RepoPrometheus:
		return &noopRepo{}
	case config.RepoVictoriaMetrics:
		return &noopRepo{}
	case config.RepoLoki:
		return &noopRepo{}
	case config.RepoElasticsearch:
		return &noopRepo{}
	default:
		return nil
	}
}

// noopRepo is a stub repository implementation.
type noopRepo struct{}

func (n *noopRepo) Query(ctx context.Context, q string) ([]byte, error) {
	return nil, nil
}
