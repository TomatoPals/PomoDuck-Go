package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_CONNECTION_DSN string `mapstructure:"DB_CONNECTION_DSN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}