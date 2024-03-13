package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := auth(c)
		c.Set("authorization", authorization)
	}
}

func auth(c *gin.Context) bool {
	_, err := c.Cookie("login")
	authorization := true
	if err != nil {
		authorization = false
	}
	return authorization
}
