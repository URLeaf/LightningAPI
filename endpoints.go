package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var serveStatic = func(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

var getLink = func(c *gin.Context) {
	link := c.PostForm("link")
	id := makeID(6)

	c.JSON(200, gin.H{"link": link, "id": id})
}
