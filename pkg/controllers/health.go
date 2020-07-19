package controllers

import (
	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context) {
	// write data and return
	c.Data(200, "text", []byte("healthy"))
}
