package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	PGUserName   string `mapstructure:"POSTGRES_USERNAME"`
	PGPassword   string `mapstructure:"POSTGRES_PASSWORD"`
	PgSSLMode    string `mapstructure:"PG_SSL_MODE"`
	PGDBmsName   string `mapstructure:"PG_DBMS"`
	PGHost       string `mapstructure:"POSTGRES_HOST"`
	PgDriverName string `mapstructure:"PG_DRIVER_NAME"`
	PGDBName     string `mapstructure:"PG_DB_NAME"`
	PgPort       string `mapstructure:"PG_PORT"`
	Host         string `mapstructure:"HOST"`
	ServerPort   string `mapstructure:"SERVER_PORT"`
}

var envs = []string{
	"POSTGRES_USERNAME", "POSTGRES_PASSWORD", "PG_SSL_MODE", "PG_DBMS", "POSTGRES_HOST",
	"PG_DRIVER_NAME", "PG_PORT", "HOST", "SERVER_PORT",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigFile("./pkg/config/loadev/.env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
