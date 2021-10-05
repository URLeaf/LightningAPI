package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// Cors
	app.Use(cors.Default())

	// Serving static
	app.Use(static.Serve("/", static.LocalFile("static/", false)))
	app.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "static/index.html", gin.H{})
	})

	// Database
	mongodbInit()

	// Routes
	app.POST("/api", getLink)
	app.GET("/get/:id", findLink)
	app.GET("/:id", hrefMe)

	// Run
	app.Run(":80")

	defer client.Disconnect(ctx)
}
