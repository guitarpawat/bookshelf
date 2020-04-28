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
	r.FuncMap = getFuncMap()
	r.LoadHTMLGlob("html/*.gohtml")
	r.NoRoute(handleNotFound)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		books, _, err := repo.GetBooksRepo().GetPaginationSortByTimeDesc(3, "")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Home | My Bookshelf",
			"books": books,
		})
	})
	r.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add", gin.H{
			"title": "Add Book | My Bookshelf",
		})
	})
	r.GET("/add-book", func(c *gin.Context) {
		err := repo.GetBooksRepo().Save(dto.Book{
			ID:      "",
			Title:   "Test",
			Edition: "1",
			Author:  []string{"GuITaRPaWaT"},
			Tags:    []string{"test", "hello"},
			Type:    dto.SoftCover,
			Status:  dto.Read,
			Volume:  nil,
			AddTime: time.Now(),
		})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Next()
	})
	r.GET("/get", func(c *gin.Context) {
		book, _, err := repo.GetBooksRepo().GetPaginationSortByTimeDesc(5, "")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(200, book)
		c.Next()
	})

	log.Fatalln(r.Run(":" + strconv.Itoa(port)))
}
