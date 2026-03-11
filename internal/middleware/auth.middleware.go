package middleware

import (
	"github.com/duongha/go-ecommerce/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.Error(c, 401, "Unauthorized! invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
