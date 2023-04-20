package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func ReadConfigFile(cfg string) (err error) {
	viper.AddConfigPath("./")
	viper.AddConfigPath("./config")
	cfg = strings.Replace(cfg, ".json", "", 1)
	viper.SetConfigName(cfg)
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	fmt.Printf("all settings: %v\n", viper.AllSettings())
	return
}
