package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	r := gin.Default()

	// load static assets
	r.Static("/public", "./public")

	// load HTML templates
	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello World",
		})
	})

	r.Run()
}
