// Package controllers provides ...
package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/yuansmin/health-recoder/pkg/models"
)

type Exercise struct{}

func (c *Exercise) List(ctx *gin.Context) {
	// todo: authorize user
	user := models.User{}
	user.ID = 1
	exerciseList, err := models.ListUserExerciseRecord(&user)
	if err != nil {
		// 500
		ctx.JSON(500, newApiError(InternalErr, err.Error()))
		return
	}

	// write data and return
	ctx.JSON(200, exerciseList)
}

func (c *Exercise) Create(ctx *gin.Context) {
	// todo: authorize user
	user := models.User{}
	user.ID = 1
	exercise := models.ExerciseRecord{}
	if err := ctx.ShouldBind(&exercise); err != nil {
		ctx.AbortWithStatusJSON(400, newApiError(BadRequestErr, err.Error()))
		return
	}

	exercise.UserID = user.ID
	if err := models.CreateExerciseRecord(&exercise); err != nil {
		ctx.AbortWithStatusJSON(500, newApiError(InternalErr, err.Error()))
		return
	}
	ctx.JSON(200, &exercise)
}

func (c *Exercise) Get(ctx *gin.Context) {
	idRaw := ctx.Param("id")
	id, err := strconv.ParseUint(idRaw, 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(400, newApiError(BadRequestErr, fmt.Sprintf("bad exercise id: %s", idRaw)))
		return
	}

	exercise := models.ExerciseRecord{}
	exercise.ID = uint(id)
	if err := models.GetExerciseRecord(&exercise); err != nil {
		ctx.AbortWithStatusJSON(500, newApiError(InternalErr, err.Error()))
		return
	}

	ctx.JSON(200, exercise)
}
