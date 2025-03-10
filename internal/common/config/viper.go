package config

import (
	"github.com/spf13/viper"
)

func NewViperConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../common/config")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}
