package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() (r *gin.Engine) {
	r = gin.Default()
	r.GET("/ping", showJsonPong)
	return
}
