package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guitarpawat/bookshelf/db"
	"github.com/guitarpawat/bookshelf/dto"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Listen(port int, repo db.Factory) {
	r := gin.Default()
	r.LoadHTMLGlob("html/*.gohtml")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Hello Home!",
		})
	})
	r.GET("/add", func(c *gin.Context) {
		err := repo.GetBooksRepo().Save(dto.Book{
			ID:      "",
			Title:   "Test",
			Edition: "1",
			Author:  []string{"GuITaRPaWaT"},
			Tags:    []string{"test", "hello"},
			Type:    dto.SoftCover,
			Status:  dto.Read,
			Volume:  nil,
			Owner:   "GuITaRPaWaT",
			AddTime: time.Now(),
		})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Next()
	})
	r.GET("/get", func(c *gin.Context) {
		book, err := repo.GetBooksRepo().GetById("5e8d42c0d8e5fb1790501d08")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(200, book)
		c.Next()
	})
	log.Fatalln(r.Run(":" + strconv.Itoa(port)))
}
