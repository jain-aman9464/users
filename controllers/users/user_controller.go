package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
