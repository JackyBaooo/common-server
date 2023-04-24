package main

import (
	"common-server/config"
	"common-server/controller"
	"common-server/model"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfgTmp := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()
	if *version {
		fmt.Println(config.VERSION)
		os.Exit(0)
	}
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	err := config.ReadConfigFile(*cfgTmp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// init log
	err = config.InitLog()
	if err != nil {
		log.Fatal(err)
	}
	// init global config
	err = config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	// init db
	err = model.InitMongoDriver()
	if err != nil {
		log.Fatal(err)
	}
	// todo init redis
	// start server
	go controller.StartServer()
	// keep server running
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Info("exit server...goodbye")
		os.Exit(0)
	}()
	select {}
}
