package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	ConsumerKey    string
	ConsumerSecret string
	Token          string
	TokenSecret    string
}

type DatabaseConfig struct {
	Username string `envconfig:"DB_USERNAME" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

func (d *DatabaseConfig) GetDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
	)
}

func ReadDBConfig() (*DatabaseConfig, error) {
	var config DatabaseConfig
	if err := envconfig.Process("bot", &config); err != nil {
		return nil, fmt.Errorf("failed to read dsn from env")
	}
	return &config, nil
}
