package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	GlobalViper  *viper.Viper
	GlobalConfig Configuration
)

type Configuration struct {
}

func InitConfig() (err error) {
	GlobalViper = viper.GetViper()
	b, err := json.Marshal(GlobalViper.AllSettings())
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("GlobalViper: %s", string(b))
	// todo set GlobalConfig
	log.Infof("GlobalConfig: %+v", GlobalConfig)
	log.Infof("InitConfig success")
	return nil
}
