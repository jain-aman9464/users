package users

import (
	"net/http"
	"strconv"

	"github.com/federicoleon/users/domain/users"
	"github.com/federicoleon/users/services"
	"github.com/federicoleon/users/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user users.User
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		// Handle error
		restErr := errors.NewBadRequestError("userId should be a number")
		c.JSON(restErr.Code, restErr)
		return
	}
	result, saveErr := services.GetUser(userID)
	if saveErr != nil {
		// handle error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		// Handle error
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// handle error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
