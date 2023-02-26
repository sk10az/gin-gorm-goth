package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Auth routes",
		})
	})

	auth.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Auth routes",
		})
	})
}