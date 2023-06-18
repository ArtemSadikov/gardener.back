package config

import "github.com/spf13/viper"

type Config struct {
	APP struct {
		PORT int `mapstructure:"app.port"`
	} `mapstructure:"app"`
}

func UnmarshalTo(path string, name string, cfg interface{}) error {
	viper.SetConfigFile(path)

	viper.SetConfigName(name)

	viper.AutomaticEnv()
	res := Config{}

	if err := viper.Unmarshal(&res); err != nil {
		return err
	}

	return nil
}
