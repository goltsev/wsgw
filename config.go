package main

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Key     string
	Secret  string
	URL     string
	Expires time.Time
}

func ReadConfigViper() (*Config, error) {
	viper.SetDefault("key", "")
	viper.SetDefault("secret", "")
	viper.SetDefault("url", "")
	viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Config{
		Key:     viper.GetString("key"),
		Secret:  viper.GetString("secret"),
		URL:     viper.GetString("url"),
		Expires: viper.GetTime("expires"),
	}, nil
}
