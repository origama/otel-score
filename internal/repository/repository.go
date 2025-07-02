package repository

import "context"

// Repository is a generic interface to query telemetry data from a
// remote observability backend. Actual implementations may query
// Prometheus, Loki, Elasticsearch, etc.
type Repository interface {
	// Query executes the provided query string and returns raw bytes
	// with the results. The format of the query and results depends
	// on the backend implementation.
	Query(ctx context.Context, q string) ([]byte, error)
}
