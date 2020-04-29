package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guitarpawat/bookshelf/db"
	"github.com/guitarpawat/bookshelf/dto"
	"github.com/guitarpawat/bookshelf/util"
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
			"title":        "Add Book | My Bookshelf",
			"bookTypes":    dto.GetBookTypes(),
			"bookStatuses": dto.GetBookStatuses(),
			"state":        c.Query("state"),
		})
	})
	r.POST("/add", func(c *gin.Context) {
		bookType, err := dto.ToBookType(util.MustGetString(c.GetPostForm("type")))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		bookStatus, err := dto.ToBookStatus(util.MustGetString(c.GetPostForm("status")))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		book := dto.Book{
			ID:      "",
			Title:   util.MustGetString(c.GetPostForm("title")),
			Edition: util.MustGetString(c.GetPostForm("edition")),
			Author:  util.SeparateComma(util.MustGetString(c.GetPostForm("authors"))),
			Tags:    util.SeparateComma(util.MustGetString(c.GetPostForm("tags"))),
			Type:    bookType,
			Status:  bookStatus,
			Volume:  util.SeparateComma(util.MustGetString(c.GetPostForm("volumes"))),
			AddTime: time.Now(),
		}

		_, err = repo.GetBooksRepo().Save(book)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Redirect(http.StatusSeeOther, "/add?state=add-success")
	})

	log.Fatalln(r.Run(":" + strconv.Itoa(port)))
}
