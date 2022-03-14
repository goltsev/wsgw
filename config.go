package main

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Key     string
	Secret  string
	URL     string
	Expires time.Time
}

func ReadConfigEnv() (*Config, error) {
	exp, err := time.Parse(time.RFC3339, os.Getenv("EXPIRES"))
	if err != nil {
		return nil, err
	}
	return &Config{
		Key:     os.Getenv("KEY"),
		Secret:  os.Getenv("SECRET"),
		URL:     os.Getenv("URL"),
		Expires: exp,
	}, nil
}

func ReadConfigViper() *Config {
	viper.SetDefault("key", "")
	viper.SetDefault("secret", "")
	viper.SetDefault("url", "")
	viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	return &Config{
		Key:     viper.GetString("key"),
		Secret:  viper.GetString("secret"),
		URL:     viper.GetString("url"),
		Expires: viper.GetTime("expires"),
	}
}
