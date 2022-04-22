package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
	dao2 "github.com/yuansmin/health-recoder/pkg/dao"

	"github.com/yuansmin/health-recoder/pkg/router"
)

var (
	dbURL = flag.String("db", "./data.db", "db url, eg: /Users/fancy/go/src/github.com/yuansmin/health-recoder/data.db")
)

func main() {
	flag.Parse()

	initLogger()
	// todo: a better way
	dao, err := dao2.New(*dbURL)
	if err != nil {
		log.Fatalf("new dao err: %s", err)
	}

	r := router.New(dao)
	r.Run()
}

func initLogger() {
	log.SetLevel(log.DebugLevel)
}
