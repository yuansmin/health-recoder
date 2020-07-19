// this is api function tests
package tests

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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
