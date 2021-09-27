package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("static/*")

	app.GET("/", serveStatic)
	app.POST("/api", getLink)

	app.Run(":2000")
}
