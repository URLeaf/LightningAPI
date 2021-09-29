package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var serveStatic = func(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

var getLink = func(c *gin.Context) {
	var req LinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorr": err.Error(),
		})
		return
	}

	fmt.Println(req.Link)

	id := makeID(6)

	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"link": req.Link,
	})

}
