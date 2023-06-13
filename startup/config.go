package startup

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	GlobalViper  *viper.Viper
	GlobalConfig *Configuration
)

type Configuration struct {
	Mongo
}

type Mongo struct {
	Url     string
	DB      string
	Timeout int
}

func defaultConfiguration() *Configuration {
	return &Configuration{
		Mongo{
			Url:     "mongodb://localhost",
			DB:      "test",
			Timeout: 5,
		},
	}
}

func setGlobalConfig() *Configuration {
	config := defaultConfiguration()
	if GlobalViper.InConfig("mongodb.url") {
		config.Url = GlobalViper.GetString("mongodb.url")
	}
	if GlobalViper.InConfig("mongodb.db") {
		config.DB = GlobalViper.GetString("mongodb.db")
	}
	if GlobalViper.InConfig("mongodb.timeout") {
		config.Timeout = GlobalViper.GetInt("mongodb.timeout")
	}
	return config
}

func InitConfig() (err error) {
	GlobalViper = viper.GetViper()
	b, err := json.MarshalIndent(GlobalViper.AllSettings(), "", "  ")
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("GlobalViper:\n%s", string(b))
	// set GlobalConfig
	GlobalConfig = setGlobalConfig()
	log.Infof("GlobalConfig: %+v", GlobalConfig)
	log.Infof("InitConfig success")
	return nil
}
