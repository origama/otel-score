package plugins

import (
	"context"

	"otelscore/internal/repository"
)

// ImpactLevel represents the severity level of a rule.
type ImpactLevel string

const (
	ImpactCritical  ImpactLevel = "Critical"
	ImpactImportant ImpactLevel = "Important"
	ImpactNormal    ImpactLevel = "Normal"
	ImpactLow       ImpactLevel = "Low"
)

// Result contains the outcome of evaluating a rule.
type Result struct {
	ID          string
	Description string
	Impact      ImpactLevel
	Passed      bool
	Error       error
}

// RuleChecker defines the interface implemented by rule plugins.
type RuleChecker interface {
	ID() string
	Description() string
	Impact() ImpactLevel
	// Check evaluates the rule using telemetry obtained via the provided
	// Repository Manager.
	Check(ctx context.Context, mgr *repository.Manager) (bool, error)
}
