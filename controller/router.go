package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() (r *gin.Engine) {
	r = gin.Default()
	r.LoadHTMLGlob("view/*")

	r.GET("/ping", showJsonPong) // sample
	r.GET("/", showHomePage)
	return
}
