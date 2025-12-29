package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	AppPort string

	LineChannelSecret      string
	LineChannelAccessToken string
	LineAdminTo            string // userId/groupId (optional)
}

func Load() (Config, error) {
	cfg := Config{
		AppPort: getenv("APP_PORT", "8080"),

		LineChannelSecret:      strings.TrimSpace(os.Getenv("LINE_CHANNEL_SECRET")),
		LineChannelAccessToken: strings.TrimSpace(os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")),
		LineAdminTo:            strings.TrimSpace(os.Getenv("LINE_ADMIN_TO")),
	}

	if cfg.LineChannelSecret == "" || cfg.LineChannelAccessToken == "" {
		return cfg, fmt.Errorf("missing LINE_CHANNEL_SECRET or LINE_CHANNEL_ACCESS_TOKEN")
	}
	return cfg, nil
}

func getenv(k, def string) string {
	v := strings.TrimSpace(os.Getenv(k))
	if v == "" {
		return def
	}
	return v
}
