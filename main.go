package main

import (
	"common-server/controller"
	"common-server/startup"
	"common-server/utils/mongo"
	"common-server/utils/redis"
	"flag"
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
	startup.HandleFlag(cfgTmp, version, help)
	// init log
	err := startup.InitLog()
	if err != nil {
		log.Fatal(err)
	}
	// init global config
	err = startup.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	// init mongo
	err = mongo.InitMongo()
	if err != nil {
		log.Fatal(err)
	}
	// init redis
	err = redis.InitRedis()
	if err != nil {
		log.Fatal(err)
	}
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
