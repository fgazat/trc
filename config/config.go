package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
)

type Config struct {
	Token       string `config:"TRACKER_TOKEN,required,backend=env"`
	XCloudOrgID string `config:"X_CLOUD_ORG_ID,backend=env"`
	XOrgID      string `config:"X_ORG_ID,backend=env"`

	APIBaseURL string `yaml:"api_base_url"`
	WebBaseURL string `yaml:"web_base_url"`

	Issues struct {
		DefaultQueue string `yaml:"default_queue"`
		Assignee     string
	}
	Filters []Filter
	Debug   bool
	Force   bool
}

type Filter struct {
	Name  string
	Query string
}

func Init() (*Config, error) {
	cfg := &Config{
		APIBaseURL: "https://api.tracker.yandex.net",
		WebBaseURL: "https://tracker.yandex.com",
		Filters: []Filter{
			{
				Name:  "My issues",
				Query: "Assignee: me()",
			},
		},
	}
	backends := []backend.Backend{
		env.NewBackend(),
	}
	cfgPath := os.Getenv("TRC_CFG_PATH")
	if cfgPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("can't get user home dir: %v", err)
		}
		cfgPath = filepath.Join(home, ".trc", "config.yaml")
	}
	if _, err := os.Stat(cfgPath); !os.IsNotExist(err) {
		backends = append(backends, file.NewBackend(cfgPath))
	}
	if err := confita.NewLoader(backends...).Load(context.Background(), cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
