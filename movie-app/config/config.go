package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

func Load() (*Config, error) {
	viper.AddConfigPath("config")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	var cfg Config

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.UnmarshalExact(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}