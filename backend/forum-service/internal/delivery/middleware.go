package delivery

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	logger *log.Logger
}

func NewAuthMiddleware(logger *log.Logger) *AuthMiddleware {
	return &AuthMiddleware{logger: logger}
}

func (m *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			return
		}
		c.Next()
	}
}

func LoggerMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Printf("%s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}
