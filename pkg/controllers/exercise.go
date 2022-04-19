// Package controllers provides ...
package controllers

import (
	"fmt"
	"strconv"

	"github.com/yuansmin/health-recoder/pkg/dao"

	"github.com/gin-gonic/gin"

	"github.com/yuansmin/health-recoder/pkg/models"
)

type exercise struct {
	dao dao.Dao
}

func NewExercise(dao dao.Dao) *exercise {
	return &exercise{dao: dao}
}

type ListExerciseRecordRequest struct {
	Offset int `json:"offset" binding:"-"`
	Limit  int `json:"limit" binding:"-"`
}

func (c *exercise) List(ctx *gin.Context) {
	// todo: authorize user
	var req ListExerciseRecordRequest
	var err error
	if err = ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, newApiError(CodeBadRequestErr, err.Error()))
		return
	}

	if err = checkPageParameter(req.Offset, req.Limit); err != nil {
		ctx.JSON(400, err)
		return
	}

	erList, err := c.dao.ExerciseRecord().List(1, req.Offset, req.Limit)
	if err != nil {
		ctx.JSON(500, newApiError(CodeInternalErr, err.Error()))
		return
	}

	// write data and return
	ctx.JSON(200, erList)
}

func (c *exercise) Create(ctx *gin.Context) {
	// todo: authorize user
	userID := 1
	exercise := models.ExerciseRecord{}
	var err error
	if err = ctx.ShouldBind(&exercise); err != nil {
		ctx.AbortWithStatusJSON(400, newApiError(CodeBadRequestErr, err.Error()))
		return
	}

	exercise.UserID = uint(userID)
	if err = c.dao.ExerciseRecord().Create(&exercise); err != nil {
		ctx.JSON(500, newApiError(CodeInternalErr, err.Error()))
		return
	}

	ctx.JSON(200, &exercise)
}

func (c *exercise) Get(ctx *gin.Context) {
	idRaw := ctx.Param("id")
	id, err := strconv.ParseUint(idRaw, 10, 32)
	if err != nil {
		ctx.JSON(400, newApiError(CodeBadRequestErr, fmt.Sprintf("bad exercise record id: %s", idRaw)))
		return
	}

	var er *models.ExerciseRecord
	if er, err = c.dao.ExerciseRecord().Get(0, uint(id)); err != nil {
		ctx.AbortWithStatusJSON(500, newApiError(CodeInternalErr, err.Error()))
		return
	}

	ctx.JSON(200, er)
}
