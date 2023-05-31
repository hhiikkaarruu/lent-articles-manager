package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showJsonPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func showHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
