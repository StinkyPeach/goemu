package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    "ok",
		"message": "ok",
	})
}
