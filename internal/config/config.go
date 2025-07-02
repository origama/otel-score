package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

// RepoKind specifies the type of remote observability repository.
type RepoKind string

const (
	RepoPrometheus      RepoKind = "prometheus"
	RepoVictoriaMetrics RepoKind = "victoriametrics"
	RepoLoki            RepoKind = "loki"
	RepoElasticsearch   RepoKind = "elasticsearch"
)

// RepositoryConfig contains connection info for a repository.
type RepositoryConfig struct {
	Name     string   `yaml:"name"`
	Type     RepoKind `yaml:"type"`
	Endpoint string   `yaml:"endpoint"`
	Username string   `yaml:"username,omitempty"`
	Password string   `yaml:"password,omitempty"`
}

// Config holds configuration values for the scoring tool.
type Config struct {
	Repositories []RepositoryConfig `yaml:"repositories"`
}

// Load reads configuration from a YAML file whose path is defined by
// the OTEL_SCORE_CONFIG_FILE environment variable. If the file is not
// found or cannot be parsed, an empty configuration is returned.
func Load() Config {
	path := os.Getenv("OTEL_SCORE_CONFIG_FILE")
	if path == "" {
		path = "config.yaml"
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}
	}
	return cfg
}
