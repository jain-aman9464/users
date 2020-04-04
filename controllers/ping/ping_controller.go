package ping

import "github.com/gin-gonic/gin"

import "net/http"

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
