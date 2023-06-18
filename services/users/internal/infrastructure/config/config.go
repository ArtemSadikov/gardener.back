package config

import (
	"github.com/spf13/viper"
)

func New() (*Config, error) {
	result := Config{}

	viper.AddConfigPath("config")

	viper.SetConfigName("default")

	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
