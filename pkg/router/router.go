package router

import (
	"github.com/gin-gonic/gin"

	"github.com/yuansmin/health-recoder/pkg/controllers"
	"github.com/yuansmin/health-recoder/pkg/dao"
)

func New(dao *dao.Dao) *gin.Engine {
	router := gin.Default()

	router.GET("/healthz", controllers.Healthz)

	// todo: config api prefix
	userController := &controllers.User{}
	router.GET("/api/users", userController.List)
	router.POST("/api/users", userController.Create)
	router.GET("/api/users/:id", userController.Get)
	router.DELETE("/api/users/:id", userController.Delete)

	exerciseController := controllers.NewExercise(dao)
	// todo: record?
	router.GET("/api/exercises", exerciseController.List)
	router.POST("/api/exercises", exerciseController.Create)
	router.GET("/api/exercises/:id", exerciseController.Get)

	return router
}
