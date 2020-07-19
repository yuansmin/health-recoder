package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/yuansmin/health-recoder/pkg/router"
)

func main() {
	initLogger()
	// todo: a better way
	r := router.RegisterAllRoutes()

	r.Run()
}

func initLogger() {
	log.SetLevel(log.DebugLevel)
}
