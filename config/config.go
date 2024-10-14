package config

import (
	"context"
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
	Debug bool
}

func Init() (*Config, error) {
	cfg := &Config{
		APIBaseURL: "https://api.tracker.yandex.net",
		WebBaseURL: "https://tracker.yandex.com",
	}
	backends := []backend.Backend{
		env.NewBackend(),
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	cfgPath := filepath.Join(home, ".trc", "config.yaml")
	if _, err = os.Stat(cfgPath); !os.IsNotExist(err) {
		backends = append(backends, file.NewBackend(cfgPath))
	}
	if err = confita.NewLoader(backends...).Load(context.Background(), cfg); err != nil {
		return nil, err
	}
	// log.Printf("%+v", *cfg)
	return cfg, nil
}
