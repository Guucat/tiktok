package mid

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context, msg string, data gin.H) {
	h := gin.H{
		"status_code": 0,
		"status_msg":  msg,
	}
	for k, v := range data {
		h[k] = v
	}
	c.JSON(http.StatusOK, h)
}

func Fail(c *gin.Context, msg string, data gin.H) {
	h := gin.H{
		"status_code": 1,
		"status_msg":  msg,
	}
	for k, v := range data {
		h[k] = v
	}
	c.JSON(http.StatusOK, h)
}
