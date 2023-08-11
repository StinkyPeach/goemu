package routers

import (
	"github.com/gin-gonic/gin"
	"goemu/api/v1"
)

func InitRouter() {
	r := gin.New()

	r.GET("/api/v1/test", v1.TestView)

	r.Run("0.0.0.0:8000")
}
