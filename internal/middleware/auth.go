package middleware

import (
	"github.com/gin-gonic/gin"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
