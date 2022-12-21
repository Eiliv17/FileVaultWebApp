package main

import (
	"net/http"

	"github.com/Eiliv17/FileVaultWebApp/controllers"
	"github.com/Eiliv17/FileVaultWebApp/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.CreateFilesDir()
}

func main() {
	r := gin.Default()

	// load static assets
	r.Static("/public", "./public")

	// load HTML templates
	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	rfile := r.Group("/file")
	{
		rfile.POST("/upload", controllers.Upload)
	}

	r.Run()
}
