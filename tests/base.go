// this is api function tests
package tests

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/yuansmin/health-recoder/pkg/models"

	dao2 "github.com/yuansmin/health-recoder/pkg/dao"
	"github.com/yuansmin/health-recoder/pkg/router"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

var (
	endpoint = "http://127.0.0.1:8080"

	client *http.Client
)

func init() {
	if ep, ok := os.LookupEnv("ENDPOINT"); ok {
		endpoint = ep
	} else {
		log.Infof("use default endpoint %s, you can change it by set environment: ENDPOINT", endpoint)
	}
	endpoint = strings.TrimRight(endpoint, "/")

	client = http.DefaultClient
	client.Timeout = time.Second * 6
}

// path: eg: /v1/users
func genURL(path string) string {
	return fmt.Sprintf("%s%s", endpoint, path)
}

func setupRouter() (r *gin.Engine, dao *dao2.Dao, clean func()) {
	rand.Seed(time.Now().Unix())

	seed := rand.Int()
	dbFile := fmt.Sprintf("/tmp/test-db-%d", seed)
	f, err := os.Create(dbFile)
	if err != nil {
		log.Fatalf("init db file err: %s", err)
	}
	f.Close()

	if err = models.AutoMigrate(dbFile); err != nil {
		log.Fatalf("AutoMigrate err: %s", err)
	}

	dao, err = dao2.New(dbFile)
	if err != nil {
		log.Fatalf("new dao err: %s", err)
	}
	clean = func() {
		os.RemoveAll(dbFile)
	}

	r = router.New(dao)
	return r, dao, clean
}
