package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("static/*")

	mongodbInit()

	app.GET("/", serveStatic)
	app.POST("/api", getLink)
	app.GET("/link/:id", hrefMe)

	app.Run(":4000")

	defer client.Disconnect(ctx)
}
