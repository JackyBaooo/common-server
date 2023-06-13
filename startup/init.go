package startup

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// VERSION change log
const VERSION = "v0.0.1"

func HandleFlag(cfgTmp *string, version, help *bool) {
	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	err := readConfigFile(*cfgTmp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readConfigFile(cfg string) (err error) {
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
