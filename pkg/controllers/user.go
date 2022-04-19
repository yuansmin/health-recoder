package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/yuansmin/health-recoder/pkg/models"
)

type User struct{}

func (u *User) List(c *gin.Context) {
	users, err := models.ListUsers()
	if err != nil {
		// 500
		c.JSON(500, newApiError(CodeInternalErr, err.Error()))
		return
	}

	// write data and return
	c.JSON(200, users)
}

func (u *User) Create(c *gin.Context) {
	user := models.User{}
	// d, _ := c.GetRawData()
	// log.Debugf("request body: %s", d)
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(400, newApiError(CodeBadRequestErr, err.Error()))
		return
	}
	if err := models.CreateUser(&user); err != nil {
		c.AbortWithStatusJSON(500, newApiError(CodeInternalErr, err.Error()))
		return
	}

	c.JSON(201, user)
}
func (u *User) Get(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.ParseUint(idRaw, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, newApiError(CodeBadRequestErr, fmt.Sprintf("bad user id: %s", idRaw)))
		return
	}
	user := models.User{}
	user.ID = uint(id)
	// todo: handle 404 error
	if err := models.GetUser(&user); err != nil {
		c.JSON(500, newApiError(CodeInternalErr, err.Error()))
		return
	}

	c.JSON(200, &user)
}

func (u *User) Delete(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := strconv.ParseUint(idRaw, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(400, newApiError(CodeBadRequestErr, fmt.Sprintf("bad user id: %s", idRaw)))
		return
	}

	user := models.User{}
	user.ID = uint(id)
	if err := models.DeleteUser(&user); err != nil {
		c.AbortWithStatusJSON(500, newApiError(CodeInternalErr, err.Error()))
		return
	}

	c.JSON(200, user)
}
