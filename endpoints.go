package main

import (
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

	var res = LinkResponse{
		ID:   makeID(6),
		Link: req.Link,
	}

	PostLink(res)
	c.JSON(http.StatusOK, res)
}

var hrefMe = func(c *gin.Context) {
	id := c.Param("id")
	takeMe(id)

	c.String(http.StatusOK, id)
}
