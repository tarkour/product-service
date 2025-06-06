package config

import (
	"fmt"
)

type DatabaseConfig struct {
	Host     string `yaml: "host"`
	Port     int    `yaml: "port"`
	User     string `yaml: "user"`
	Password string `yaml: "password"`
	DBName   string `yaml: "dbname"`
	SSLMode  string `yaml: "sslmode"`
}

type TelegramConfig struct {
	token     string `yaml: "token"`
	admin_id  int64  `yaml: "admin_id"`
	safe_mode bool   `yaml: "safe_mode"`
}

type Config struct {
	Database DatabaseConfig `yaml: "database"`
	Telegram TelegramConfig `yaml: "telegram"`
}

func (d *DatabaseConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DBName,
		d.SSLMode,
	)
}
