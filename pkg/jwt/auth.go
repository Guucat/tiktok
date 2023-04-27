package jwt

import (
	"github.com/gin-gonic/gin"
	"tiktok/mid"
)

// Auth the authentication middleware
func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}

		mes, err := ParseToken(token)
		if err != nil {
			mid.Fail(c, "invalid token", nil)
			c.Abort()
			return
		}

		c.Set("id", mes.Id)
		c.Next()
	}
}
