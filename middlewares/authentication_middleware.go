package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthenticationMiddleware(c *gin.Context) {
	if false {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()

	log.Println("Hi")
}
