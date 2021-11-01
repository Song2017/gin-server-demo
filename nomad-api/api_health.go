package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth - health
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
