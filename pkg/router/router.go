package router

import (
	"github.com/gin-gonic/gin"

	"github.com/yuansmin/health-recoder/pkg/controllers"
)

func RegisterAllRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/healthz", controllers.Healthz)

	userController := &controllers.User{}
	router.GET("/api/users", userController.List)
	router.POST("/api/users", userController.Create)
	router.GET("/api/users/:id", userController.Get)
	router.DELETE("/api/users/:id", userController.Delete)

	return router
}
