package config

import (
	"errors"

	"github.com/spf13/viper"
)

var errEmptyToken = errors.New("empty token. Please, export TG_TOKEN to env")

type Messages struct {
	Responses
	Errors
}

type Config struct {
	TelegramToken  string
	PollingTimeout int64  `mapstructure:"polling_timeout"`
	DataBaseFile   string `mapstructure:"db_file"`
	Messages       Messages
}

type Responses struct {
	Start string `mapstructure:"start"`
}

type Errors struct {
	Default string `mapstructure:"default"`
}

func LoadConfig() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}

	if cfg.TelegramToken == "" {
		return nil, errEmptyToken
	}
	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.error", &cfg.Messages.Errors); err != nil {
		return err
	}

	return nil
}

const env_tg_token = "TGTOKEN"

func fromEnv(cfg *Config) error {
	if err := viper.BindEnv(env_tg_token); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString(env_tg_token)

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
