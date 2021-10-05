package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

var serveStatic = func(c *gin.Context) {
	c.HTML(http.StatusOK, "static/index.html", gin.H{})
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

var findLink = func(c *gin.Context) {
	id := c.Param("id")
	mylink := takeMe(id)

	c.JSON(http.StatusOK, mylink)
}

var hrefMe = func(c *gin.Context) {
	id := c.Param("id")
	mylink := fmt.Sprintf("%v", takeMe(id))
	location := url.URL{Path: "//" + mylink}
	c.Redirect(http.StatusFound, location.RequestURI())
}
