package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLog() (err error) {
	switch viper.GetString("log.level") {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		err = fmt.Errorf("unknown log level: %s", viper.GetString("log.level"))
		return err
	}
	log.Infof("InitLog success: %s", viper.GetString("log.level"))
	return nil
}
