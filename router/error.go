package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404", gin.H{
		"title": "Not Found | My Bookshelf",
	})
}
