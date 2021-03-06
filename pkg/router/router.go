package router

import (
	"github.com/gin-gonic/gin"

	"github.com/yuansmin/health-recoder/pkg/controllers"
)

func RegisterAllRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/healthz", controllers.Healthz)

	// todo: config api prefix
	userController := &controllers.User{}
	router.GET("/api/users", userController.List)
	router.POST("/api/users", userController.Create)
	router.GET("/api/users/:id", userController.Get)
	router.DELETE("/api/users/:id", userController.Delete)

	exerciseController := &controllers.Exercise{}
	router.GET("/api/exercise", exerciseController.List)
	router.POST("/api/exercise", exerciseController.Create)
	router.GET("/api/exercise/:id", exerciseController.Get)

	return router
}
